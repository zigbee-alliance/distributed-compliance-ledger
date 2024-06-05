#!/usr/bin/env bash
# parts are copied from https://github.com/cosmos/cosmos-sdk/blob/master/scripts/protoc-swagger-gen.sh

# read swagger config and output paths from cmd
CONFIG_FILE=${1:-''}
OUTPUT_FILE=${2:-''}

set -eo pipefail

mkdir -p ./swagger-out

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do

  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then
    buf generate --template proto/buf.gen.swagger.yaml "$query_file"
    cp -r zigbeealliance/distributedcomplianceledger/* swagger-out/
  fi
done

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
pushd ./swagger-out
swagger-combine $CONFIG_FILE -o $OUTPUT_FILE --continueOnConflictingPaths true --includeDefinitions true
popd

rm -rf zigbeealliance
rm -rf ./swagger-out