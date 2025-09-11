#!/bin/bash

set -euo pipefail

# Input file
COVERAGE_FILE="cover.out"

# Temporary file
TEMP_FILE="cover.out.tmp"

# Patterns to exclude
EXCLUDE_PATTERNS=(
    "/integration_tests/"
    "/tests/"
    "/testutil/"
    "/simulation/"
    "module_simulation"
    ".pb."
)

# Function to exclude lines based on patterns
exclude_lines() {
  local input_file="$1"
  local output_file="$2"

  # Read the input file line by line
  while IFS= read -r line; do
    exclude=false

    # Check if the line matches any of the exclude patterns
    for pattern in "${EXCLUDE_PATTERNS[@]}"; do
      if [[ "$line" == *$pattern* ]]; then
        exclude=true
        break
      fi
    done

    # If the line should not be excluded, write it to the output file
    if ! $exclude; then
      echo "$line" >> "$output_file"
    fi
  done < "$input_file"
}

# Exclude lines and write to the temporary file
exclude_lines "$COVERAGE_FILE" "$TEMP_FILE"

# Replace the original file with the temporary file
mv "$TEMP_FILE" "$COVERAGE_FILE"

echo "Coverage file '$COVERAGE_FILE' updated, excluding specified patterns."