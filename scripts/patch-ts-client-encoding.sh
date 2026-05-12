#!/usr/bin/env bash
#
# Patches every ts-client/*/rest.ts so axios percent-encodes query
# parameter values. Without this, base64 pagination keys (which contain
# '+', '/', '=') are mangled in transit -- '+' is decoded as space on
# the server, producing "illegal base64 data" errors on next_key follow-up
# requests. Idempotent: re-run safely after `ignite generate ts-client`.

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
TS_CLIENT_DIR="$REPO_ROOT/ts-client"

if [[ ! -d "$TS_CLIENT_DIR" ]]; then
  echo "error: $TS_CLIENT_DIR not found" >&2
  exit 1
fi

OLD_LINE='    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });'

IFS= read -r -d '' NEW_BLOCK <<'EOF' || true
    this.instance = axios.create({
      ...axiosConfig,
      baseURL: axiosConfig.baseURL || "",
      // RFC 3986 encode every value so base64 pagination keys (+, /, =) survive transit.
      paramsSerializer: (params: Record<string, any>) => {
        const parts: string[] = [];
        for (const key of Object.keys(params)) {
          const value = params[key];
          if (value === undefined || value === null) continue;
          const ek = encodeURIComponent(key);
          if (Array.isArray(value)) {
            for (const item of value) parts.push(`${ek}=${encodeURIComponent(String(item))}`);
          } else {
            parts.push(`${ek}=${encodeURIComponent(String(value))}`);
          }
        }
        return parts.join("&");
      },
    });
EOF

patched=0
skipped=0
missing=0

while IFS= read -r -d '' file; do
  if grep -q "paramsSerializer" "$file"; then
    skipped=$((skipped + 1))
    continue
  fi
  if ! grep -qF "$OLD_LINE" "$file"; then
    echo "warn: anchor line not found in $file" >&2
    missing=$((missing + 1))
    continue
  fi
  python3 - "$file" "$OLD_LINE" "$NEW_BLOCK" <<'PY'
import sys, pathlib
path = pathlib.Path(sys.argv[1])
old_line = sys.argv[2]
new_block = sys.argv[3]
text = path.read_text()
if old_line not in text:
    sys.exit(f"anchor missing in {path}")
path.write_text(text.replace(old_line, new_block, 1))
PY
  patched=$((patched + 1))
  echo "patched $file"
done < <(find "$TS_CLIENT_DIR" -type f -name "rest.ts" -print0)

echo "done. patched=$patched skipped=$skipped missing=$missing"

if [[ $missing -gt 0 ]]; then
  exit 2
fi
