#!/bin/bash
# Compare every `dcld tx ...` and `dcld query ...` subcommand invoked by each
# bash script under integration_tests/upgrade/ against the subcommands its
# migrated Go counterpart calls. Reports subcommands that bash exercises but
# Go does not — i.e. unmigrated CLI surface.
#
# Why this exists: the historical dcld binaries used by the upgrade tests are
# pre-built GitHub release artifacts and are not cover-instrumented, so the
# generic coverage-diff approach (used for CLI tests) cannot see most of the
# chain. A subcommand-level diff is the cheapest reliable parity check.
#
# Usage:
#   scripts/upgrade-command-diff.sh                       # diff + exit 0
#   scripts/upgrade-command-diff.sh --strict              # fail on new gaps
#                                                         # vs baseline
#   scripts/upgrade-command-diff.sh --update-baseline     # rewrite baseline
#
# Files:
#   upgrade_command_gap.txt                                # current gaps
#   integration_tests/upgrade/.command_baseline.txt        # accepted gaps
#
# Exit codes:
#   0  no regression (gap ⊆ baseline, or default/non-strict mode)
#   1  strict-mode regression or invocation error

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
GAP_REPORT="$REPO_ROOT/upgrade_command_gap.txt"
BASELINE="$REPO_ROOT/integration_tests/upgrade/.command_baseline.txt"

STRICT=false
UPDATE_BASELINE=false
while [[ $# -gt 0 ]]; do
  case $1 in
    --strict)           STRICT=true; shift ;;
    --update-baseline)  UPDATE_BASELINE=true; shift ;;
    -h|--help)
      sed -n '2,21p' "$0" | sed 's/^# \{0,1\}//'
      exit 0
      ;;
    *) echo "Unknown option: $1" >&2; exit 1 ;;
  esac
done

cd "$REPO_ROOT"

# pair: bash_path|go_path
PAIRS=(
  "integration_tests/upgrade/01-test-upgrade-initialize-0.12.sh|integration_tests/upgrade/migration_0_12_init_test.go"
  "integration_tests/upgrade/02-test-upgrade-0.12-rollback.sh|integration_tests/upgrade/migration_0_12_rollback_test.go"
  "integration_tests/upgrade/03-test-upgrade-0.12-to-1.2.sh|integration_tests/upgrade/migration_1_2_test.go"
  "integration_tests/upgrade/04-test-upgrade-1.2-rollback.sh|integration_tests/upgrade/migration_1_2_rollback_test.go"
  "integration_tests/upgrade/05-test-upgrade-1.2-to-1.4.3.sh|integration_tests/upgrade/migration_1_4_3_test.go"
  "integration_tests/upgrade/06-test-upgrade-1.4.3-to-1.4.4.sh|integration_tests/upgrade/migration_1_4_4_test.go"
  "integration_tests/upgrade/07-test-upgrade-1.4.4-to-1.5.1.sh|integration_tests/upgrade/migration_1_5_1_test.go"
  "integration_tests/upgrade/08-test-upgrade-1.5.1-to-1.5.2.sh|integration_tests/upgrade/migration_1_5_2_test.go"
  "integration_tests/upgrade/09-test-upgrade-1.5.2-to-1.6.0.sh|integration_tests/upgrade/migration_1_6_0_test.go"
  "integration_tests/upgrade/10-test-upgrade-1.6.0-to-master.sh|integration_tests/upgrade/migration_master_test.go"
  "integration_tests/upgrade/11-test-add-new-node-after-upgrade.sh|integration_tests/upgrade/migration_add_node_test.go"
)

# extract_bash_subcommands <file>
# Returns sorted unique "<tx|query> <module> <subcommand>" lines for every
# `dcld[_BIN_*] (tx|query) <module> <subcommand>` invocation in the bash file.
# Handles backslash-line-continuations.
extract_bash_subcommands() {
  local file="$1"
  # 1. Join continuation lines (line ending in `\` + newline → one logical line).
  # 2. Find each `dcld-like-binary  tx|query  <module>  <subcommand>` triple.
  #    The binary token can be: literal `dcld`, `./dcld`, `$DCLD_BIN[_…]`,
  #    or any other shell variable matching `\$[A-Z_]+`.
  # 3. Drop everything except the (tx|query) <module> <subcommand> tuple.
  ( awk '
    /\\$/ { sub(/\\$/, " "); buf = buf $0; next }
    { print buf $0; buf = "" }
  ' "$file" | grep -oE '(\$[A-Z_][A-Z_0-9]*|\.?/?\bdcld[a-zA-Z0-9_]*\b)[[:space:]]+(tx|query)[[:space:]]+[a-zA-Z][a-zA-Z0-9_-]*[[:space:]]+[a-zA-Z][a-zA-Z0-9_-]*' \
    | sed -E 's/^[^[:space:]]+[[:space:]]+//; s/[[:space:]]+/ /g' \
    | sort -u ) || true
}

# extract_go_subcommands <file>
# Returns sorted unique "<tx|query> <module> <subcommand>" lines reachable from
# the Go file. Catches the canonical shapes used by the migration helpers:
#   ExecuteTxWithBin(bin, "tx", "module", "subcmd", ...)
#   ExecuteCLIWithBin(bin, "query", "module", "subcmd", ...)
# and a few package-level helpers we know wrap a fixed subcommand:
#   ProposeUpgrade   → tx dclupgrade propose-upgrade
#   ApproveUpgrade   → tx dclupgrade approve-upgrade
#   RejectUpgrade    → tx dclupgrade reject-upgrade
#   QueryProposedUpgrade → query dclupgrade proposed-upgrade
#   QueryApprovedUpgrade → query dclupgrade approved-upgrade
#   QueryRejectedUpgrade → query dclupgrade rejected-upgrade
#   QueryUpgradePlan     → query upgrade plan
#   QueryAppliedPlan     → query upgrade applied
#   ProposeAddAccount    → tx auth propose-add-account
#   ApproveAddAccount    → tx auth approve-add-account
#   ProposeRevokeAccount → tx auth propose-revoke-account
#   ApproveRevokeAccount → tx auth approve-revoke-account
#   CreateKey            → keys add  (informational, not tx/query)
extract_go_subcommands() {
  local file="$1"
  {
    # Direct ExecuteTxWithBin / ExecuteCLIWithBin call shapes.
    # Calls span multiple lines, so squash newlines first.
    tr '\n' ' ' < "$file" \
      | grep -oE '"(tx|query)"[[:space:]]*,[[:space:]]*"[a-zA-Z0-9_-]+"[[:space:]]*,[[:space:]]*"[a-zA-Z0-9_-]+"' \
      | sed -E 's/"(tx|query)"[[:space:]]*,[[:space:]]*"([a-zA-Z0-9_-]+)"[[:space:]]*,[[:space:]]*"([a-zA-Z0-9_-]+)".*/\1 \2 \3/'

    # Package helpers with implicit (tx|query) module subcmd.
    grep -oE '\bProposeUpgrade\b' "$file" >/dev/null && echo "tx dclupgrade propose-upgrade"
    grep -oE '\bApproveUpgrade\b' "$file" >/dev/null && echo "tx dclupgrade approve-upgrade"
    grep -oE '\bRejectUpgrade\b'  "$file" >/dev/null && echo "tx dclupgrade reject-upgrade"
    grep -oE '\bQueryProposedUpgrade\b' "$file" >/dev/null && echo "query dclupgrade proposed-upgrade"
    grep -oE '\bQueryApprovedUpgrade\b' "$file" >/dev/null && echo "query dclupgrade approved-upgrade"
    grep -oE '\bQueryRejectedUpgrade\b' "$file" >/dev/null && echo "query dclupgrade rejected-upgrade"
    grep -oE '\bQueryUpgradePlan\b'     "$file" >/dev/null && echo "query upgrade plan"
    grep -oE '\bQueryAppliedPlan\b'     "$file" >/dev/null && echo "query upgrade applied"
    grep -oE '\bProposeAddAccount\b'    "$file" >/dev/null && echo "tx auth propose-add-account"
    grep -oE '\bApproveAddAccount\b'    "$file" >/dev/null && echo "tx auth approve-add-account"
    grep -oE '\bProposeRevokeAccount\b' "$file" >/dev/null && echo "tx auth propose-revoke-account"
    grep -oE '\bApproveRevokeAccount\b' "$file" >/dev/null && echo "tx auth approve-revoke-account"
    grep -oE '\bDisableValidatorNode\b' "$file" >/dev/null && echo "tx validator disable-node"
    grep -oE '\bEnableValidatorNode\b'  "$file" >/dev/null && echo "tx validator enable-node"
    grep -oE '\bProposeDisableValidatorNode\b' "$file" >/dev/null && echo "tx validator propose-disable-node"
    grep -oE '\bApproveDisableValidatorNode\b' "$file" >/dev/null && echo "tx validator approve-disable-node"

    # AddValidatorNode → many subcommands all rolled into one helper.
    if grep -oE '\bAddValidatorNode\b' "$file" >/dev/null; then
      cat <<EOF
tx auth propose-add-account
tx auth approve-add-account
tx validator add-node
query validator node
EOF
    fi

    # RunValidatorDisableEnableFlow wraps the standard per-script
    # disable/enable/propose-disable/approve-disable sequence.
    if grep -oE '\bRunValidatorDisableEnableFlow\b' "$file" >/dev/null; then
      cat <<EOF
tx validator disable-node
tx validator enable-node
tx validator propose-disable-node
tx validator approve-disable-node
EOF
    fi

    # CreateAndApproveAccount / createVendorWithApprovals — propose+approve.
    if grep -oE '\b(CreateAndApproveAccount|createVendorWithApprovals)\b' "$file" >/dev/null; then
      cat <<EOF
tx auth propose-add-account
tx auth approve-add-account
EOF
    fi

    # Compliance triplet loop (rollback / master post-upgrade): pattern is
    #   for _, action := range []struct{...}{
    #     {"provision-model", ...},
    #     {"certify-model", ...},
    #     {"revoke-model", ...},
    #   } {
    # The string literals appear elsewhere in the file as compliance subcommand
    # tokens; emit `tx compliance <verb>` for any we find.
    for verb in provision-model certify-model revoke-model certify-revoked-model; do
      if grep -qE "\"$verb\"" "$file"; then
        echo "tx compliance $verb"
      fi
    done

    # PKI approve/reject loop pattern (script 03 root-cert ladder):
    #   for _, action := range []string{"approve-add-x509-root-cert", "reject-add-x509-root-cert"} {
    #     ExecuteTxWithBin(... "tx", "pki", action, ...)
    #   }
    # Same trick — find the literal tokens.
    for verb in approve-add-x509-root-cert reject-add-x509-root-cert approve-revoke-x509-root-cert; do
      if grep -qE "\"$verb\"" "$file"; then
        echo "tx pki $verb"
      fi
    done

    # proposeUserAccount / revokeUserAccount → auth propose+approve flows.
    if grep -oE '\bproposeUserAccount\b' "$file" >/dev/null; then
      cat <<EOF
tx auth propose-add-account
tx auth approve-add-account
EOF
    fi
    if grep -oE '\brevokeUserAccount\b' "$file" >/dev/null; then
      cat <<EOF
tx auth propose-revoke-account
tx auth approve-revoke-account
EOF
    fi

    # SoftwareUpgradeStep.Run is the propose+approve+wait+verify-applied flow.
    if grep -oE '\bSoftwareUpgradeStep\b' "$file" >/dev/null; then
      cat <<EOF
tx dclupgrade propose-upgrade
tx dclupgrade approve-upgrade
query upgrade plan
query upgrade applied
EOF
    fi
  } | sort -u
}

: > "$GAP_REPORT"

total_bash=0
total_go=0
total_gaps=0
files_with_gaps=0

printf '%-70s %5s %5s %5s\n' "PAIR" "BASH" "GO" "GAP"
printf '%-70s %5s %5s %5s\n' "----" "----" "--" "---"

for entry in "${PAIRS[@]}"; do
  bash_file="${entry%%|*}"
  go_file="${entry##*|}"

  if [[ ! -f "$bash_file" ]]; then
    echo "WARN: $bash_file missing — was the file deleted? Skipping." >&2
    continue
  fi
  if [[ ! -f "$go_file" ]]; then
    echo "WARN: $go_file missing — skipping." >&2
    continue
  fi

  bash_cmds=$(extract_bash_subcommands "$bash_file")
  go_cmds=$(extract_go_subcommands "$go_file")

  bash_count=$(printf '%s\n' "$bash_cmds" | grep -c . || true)
  go_count=$(printf '%s\n' "$go_cmds"  | grep -c . || true)

  # Subcommands in bash but not in Go.
  missing=$(comm -23 <(printf '%s\n' "$bash_cmds") <(printf '%s\n' "$go_cmds"))
  miss_count=$(printf '%s\n' "$missing" | grep -c . || true)

  pair_label="$(basename "$bash_file") -> $(basename "$go_file")"
  printf '%-70s %5d %5d %5d\n' "$pair_label" "$bash_count" "$go_count" "$miss_count"

  if (( miss_count > 0 )); then
    files_with_gaps=$((files_with_gaps + 1))
    while IFS= read -r line; do
      [[ -z "$line" ]] && continue
      printf '%s :: %s\n' "$bash_file" "$line" >> "$GAP_REPORT"
    done <<< "$missing"
  fi

  total_bash=$((total_bash + bash_count))
  total_go=$((total_go + go_count))
  total_gaps=$((total_gaps + miss_count))
done

echo
printf 'Totals: %d bash subcommands, %d Go subcommands, %d gaps across %d pairs.\n' \
  "$total_bash" "$total_go" "$total_gaps" "$files_with_gaps"

if (( total_gaps > 0 )); then
  echo
  echo "Wrote $total_gaps entries to $GAP_REPORT"
fi

# Baseline maintenance.
if $UPDATE_BASELINE; then
  if [[ -s "$GAP_REPORT" ]]; then
    sort -u "$GAP_REPORT" > "$BASELINE"
  else
    : > "$BASELINE"
  fi
  baseline_count=$(grep -c . "$BASELINE" || true)
  echo
  echo "==> Baseline updated: $BASELINE ($baseline_count entries)"
  exit 0
fi

# Strict-mode enforcement: gap must be a subset of baseline.
if $STRICT; then
  echo
  echo "==> Enforcing baseline ($BASELINE)"

  if [[ ! -f "$BASELINE" ]]; then
    echo "    baseline file does not exist — skipping enforcement"
    echo "    create it with: scripts/upgrade-command-diff.sh --update-baseline"
    exit 0
  fi

  TMP_GAP=$(mktemp)
  TMP_BASE=$(mktemp)
  trap 'rm -f "$TMP_GAP" "$TMP_BASE"' EXIT
  sort -u "$GAP_REPORT" > "$TMP_GAP"
  sort -u "$BASELINE"   > "$TMP_BASE"

  NEW=$(comm -23 "$TMP_GAP" "$TMP_BASE")
  RESOLVED=$(comm -13 "$TMP_GAP" "$TMP_BASE")
  new_count=$(printf '%s' "$NEW"      | grep -c . || true)
  resolved_count=$(printf '%s' "$RESOLVED" | grep -c . || true)

  printf '    new regression entries:    %d\n' "$new_count"
  printf '    entries resolved since baseline: %d\n' "$resolved_count"

  if (( resolved_count > 0 )); then
    echo
    echo "    Baseline can shrink. Run:"
    echo "      scripts/upgrade-command-diff.sh --update-baseline"
    echo "    and commit the smaller $BASELINE."
  fi

  if (( new_count > 0 )); then
    echo
    echo "REGRESSION: $new_count subcommand(s) exercised by bash with no Go match and not in baseline:"
    printf '%s\n' "$NEW" | head -20 | sed 's/^/  /'
    if (( new_count > 20 )); then
      printf "  ...   (%d more — see %s)\n" $(( new_count - 20 )) "$GAP_REPORT"
    fi
    exit 1
  fi

  echo "    OK: current gap is a subset of baseline."
  exit 0
fi

if (( total_gaps == 0 )); then
  echo "No gaps. Every bash subcommand has a matching Go invocation."
fi
exit 0
