#!/bin/bash
# Compare code coverage between the bash CLI test suite and the Go CLI test
# suite to find code paths exercised by bash but missing in Go — i.e. tests
# that haven't been fully migrated yet.
#
# Usage:
#   scripts/coverage-diff.sh                       # run both suites, then diff
#   scripts/coverage-diff.sh --no-run              # reuse existing cover_*.out
#   scripts/coverage-diff.sh --no-run --strict     # enforce: fail if new gaps appear vs baseline
#   scripts/coverage-diff.sh --no-run --update-baseline   # rewrite baseline from current gap
#
# Files:
#   cover_bash.out                           # textfmt profile of bash CLI suite
#   cover_go.out                             # textfmt profile of Go CLI suite
#   coverage_gap.txt                         # blocks hit by bash but not Go (sorted)
#   integration_tests/cli/.coverage_baseline.txt
#                                            # checked-in baseline of accepted gaps
#
# Exit codes:
#   0  no regression (gap ⊆ baseline, non-strict mode, or no baseline yet)
#   1  invocation error or strict-mode regression

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
COVER_DIR="$REPO_ROOT/integration_tests/gocover"
BASH_PROFILE="$REPO_ROOT/cover_bash.out"
GO_PROFILE="$REPO_ROOT/cover_go.out"
GAP_REPORT="$REPO_ROOT/coverage_gap.txt"
BASELINE="$REPO_ROOT/integration_tests/cli/.coverage_baseline.txt"

usage() {
  sed -n '2,21p' "$0" | sed 's/^# \{0,1\}//'
  exit "${1:-0}"
}

RUN=true
STRICT=false
UPDATE_BASELINE=false
while [[ $# -gt 0 ]]; do
  case $1 in
    --no-run)           RUN=false; shift ;;
    --strict)           STRICT=true; shift ;;
    --update-baseline)  UPDATE_BASELINE=true; shift ;;
    -h|--help)          usage 0 ;;
    *) echo "Unknown option: $1" >&2; usage 1 ;;
  esac
done

cd "$REPO_ROOT"

if $RUN; then
  echo "==> Running bash CLI suite with coverage (this is slow)"
  ./integration_tests/run-all.sh cli cover
  go tool covdata textfmt -i="$COVER_DIR" -o "$BASH_PROFILE"
  echo "    wrote $BASH_PROFILE"

  echo "==> Running Go CLI suite with coverage (this is slow)"
  ./integration_tests/run-all.sh cli_go cover
  go tool covdata textfmt -i="$COVER_DIR" -o "$GO_PROFILE"
  echo "    wrote $GO_PROFILE"
fi

if [[ ! -f "$BASH_PROFILE" || ! -f "$GO_PROFILE" ]]; then
  echo "ERROR: $BASH_PROFILE or $GO_PROFILE not found. Re-run without --no-run." >&2
  exit 1
fi

echo
echo "==> Computing coverage diff (blocks hit by bash but not by Go)"
echo

awk -v gap_file="$GAP_REPORT" '
  # Exclusion patterns mirror .github/workflows/filter_coverage.sh
  function is_excluded(key) {
    if (key ~ /\/integration_tests\//) return 1
    if (key ~ /\/tests\//)             return 1
    if (key ~ /\/testutil\//)          return 1
    if (key ~ /\/simulation\//)        return 1
    if (key ~ /module_simulation/)     return 1
    if (key ~ /\.pb\./)                return 1
    return 0
  }

  /^mode:/ { next }

  {
    key = $1
    count = $3 + 0
    if (is_excluded(key)) next

    if (FILENAME == ARGV[1]) {
      bash_count[key] = count
      if (count > 0) bash_covered++
    } else {
      go_count[key] = count
      if (count > 0) go_covered++
    }
  }

  END {
    gap_blocks = 0
    n = 0
    delete gap_keys

    for (key in bash_count) {
      if (bash_count[key] > 0 && (!(key in go_count) || go_count[key] == 0)) {
        split(key, parts, ":")
        file = parts[1]
        gap_by_file[file]++
        n++
        gap_keys[n] = key
        gap_blocks++
      }
    }

    printf "Bash-covered blocks:   %d\n", bash_covered + 0
    printf "Go-covered blocks:     %d\n",  go_covered + 0
    printf "Gap (bash-only):       %d blocks across %d files\n", gap_blocks, length(gap_by_file)
    print  ""

    if (gap_blocks == 0) {
      print "No coverage gaps — Go suite covers everything bash does."
      # Still write an empty gap file so downstream tooling has a consistent path
      printf "" > gap_file
      close(gap_file)
      exit 0
    }

    # Sort files by gap count desc (simple insertion sort)
    fn = 0
    for (file in gap_by_file) { fn++; files[fn] = file }
    for (i = 2; i <= fn; i++) {
      j = i
      while (j > 1 && gap_by_file[files[j-1]] < gap_by_file[files[j]]) {
        t = files[j]; files[j] = files[j-1]; files[j-1] = t
        j--
      }
    }

    print "Top files by bash-only block count:"
    top = (fn < 20) ? fn : 20
    for (i = 1; i <= top; i++) {
      printf "  %4d  %s\n", gap_by_file[files[i]], files[i]
    }
    if (fn > top) printf "  ...   (%d more files)\n", fn - top
    print ""

    # Sort gap_keys before writing so the file is deterministic / diff-friendly
    for (i = 2; i <= gap_blocks; i++) {
      j = i
      while (j > 1 && gap_keys[j-1] > gap_keys[j]) {
        t = gap_keys[j]; gap_keys[j] = gap_keys[j-1]; gap_keys[j-1] = t
        j--
      }
    }
    for (i = 1; i <= gap_blocks; i++) print gap_keys[i] > gap_file
    close(gap_file)

    printf "Full list written to %s (%d entries)\n", gap_file, gap_blocks
  }
' "$BASH_PROFILE" "$GO_PROFILE"

# Ensure gap report exists even when zero
[[ -f "$GAP_REPORT" ]] || : > "$GAP_REPORT"

# Baseline maintenance
if $UPDATE_BASELINE; then
  cp "$GAP_REPORT" "$BASELINE"
  echo
  echo "==> Baseline updated: $BASELINE ($(wc -l < "$BASELINE" | tr -d ' ') entries)"
  exit 0
fi

# Strict-mode enforcement: gap must be a subset of baseline
if $STRICT; then
  echo
  echo "==> Enforcing baseline ($BASELINE)"
  if [[ ! -f "$BASELINE" ]]; then
    echo "    baseline file does not exist — skipping enforcement"
    echo "    create it with: scripts/coverage-diff.sh --no-run --update-baseline"
    exit 0
  fi

  # Sort both for stable comm output
  TMP_GAP=$(mktemp)
  TMP_BASE=$(mktemp)
  trap 'rm -f "$TMP_GAP" "$TMP_BASE"' EXIT
  sort -u "$GAP_REPORT"  > "$TMP_GAP"
  sort -u "$BASELINE"    > "$TMP_BASE"

  # New gaps = in current gap but NOT in baseline → regression
  NEW=$(comm -23 "$TMP_GAP" "$TMP_BASE")
  # Removed gaps = in baseline but NOT in current gap → migration progress
  RESOLVED=$(comm -13 "$TMP_GAP" "$TMP_BASE")

  new_count=$(printf '%s' "$NEW"      | grep -c . || true)
  resolved_count=$(printf '%s' "$RESOLVED" | grep -c . || true)

  printf "    new regression blocks: %d\n" "$new_count"
  printf "    blocks resolved since baseline: %d\n" "$resolved_count"

  if (( resolved_count > 0 )); then
    echo
    echo "    The baseline can shrink. Run:"
    echo "      scripts/coverage-diff.sh --no-run --update-baseline"
    echo "    and commit the updated $BASELINE"
  fi

  if (( new_count > 0 )); then
    echo
    echo "REGRESSION: $new_count bash-covered blocks are not covered by Go and not in baseline:"
    printf '%s\n' "$NEW" | head -20 | sed 's/^/  /'
    if (( new_count > 20 )); then
      printf "  ...   (%d more — see coverage_gap.txt)\n" $(( new_count - 20 ))
    fi
    exit 1
  fi

  echo "    OK: current gap is a subset of baseline."
fi
