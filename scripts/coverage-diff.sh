#!/bin/bash
# Compare code coverage between the bash CLI test suite and the Go CLI test
# suite to find code paths exercised by bash but missing in Go — i.e. tests
# that haven't been fully migrated yet.
#
# Usage:
#   scripts/coverage-diff.sh              # run both suites, then diff
#   scripts/coverage-diff.sh --no-run     # reuse existing cover_bash.out / cover_go.out
#
# Output:
#   stdout                 — summary (counts, top files by gap)
#   coverage_gap.txt       — full list of bash-only blocks, one per line

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
COVER_DIR="$REPO_ROOT/integration_tests/gocover"
BASH_PROFILE="$REPO_ROOT/cover_bash.out"
GO_PROFILE="$REPO_ROOT/cover_go.out"
GAP_REPORT="$REPO_ROOT/coverage_gap.txt"

usage() {
  sed -n '2,11p' "$0" | sed 's/^# \{0,1\}//'
  exit "${1:-0}"
}

RUN=true
while [[ $# -gt 0 ]]; do
  case $1 in
    --no-run)  RUN=false; shift ;;
    -h|--help) usage 0 ;;
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
  function is_excluded(key,    p) {
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

    # Write full detailed list to file
    for (i = 1; i <= gap_blocks; i++) print gap_keys[i] > gap_file
    close(gap_file)

    printf "Full list written to %s (%d entries)\n", gap_file, gap_blocks
  }
' "$BASH_PROFILE" "$GO_PROFILE"
