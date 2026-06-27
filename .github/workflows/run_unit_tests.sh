#!/bin/bash

# Runs the Go unit tests with coverage and writes the merged profile to cover.out.
#
# The main suite runs with -tags=dev. url_liveness.go gates its HTTP logic behind
# that dev flag (and its !dev-tagged test is excluded), so it is unreachable in the
# dev run. We re-run the utils/cli package without the dev tag to collect real
# coverage for the liveness checks and merge it in, so it shows up in the coverage
# report. The non-dev package build reports the same canonical file paths as the
# dev run, so gocovmerge unions the two cleanly.

set -euo pipefail

GOCOVMERGE_VERSION="b5bfa59ec0adc420475f97f89b58045c721d761c"

# shellcheck disable=SC2046
go test -tags=dev -json -v $(go list ./... | grep -v '/integration_tests') \
    -coverprofile=./cover.out -covermode=set -coverpkg=./... 2>&1 \
    | tee /tmp/gotest.log | gotestfmt

go test -covermode=set -coverprofile=./cover_url_liveness.out \
    -coverpkg=./utils/cli/... ./utils/cli/

go install "github.com/wadey/gocovmerge@${GOCOVMERGE_VERSION}"
gocovmerge ./cover.out ./cover_url_liveness.out > ./cover_merged.out
mv ./cover_merged.out ./cover.out
