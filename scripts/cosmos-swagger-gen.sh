#!/usr/bin/env bash
# possible values: "base" or "tx"
TYPE=${1:-tx}

echo "Generating Cosmos openapi for '$TYPE'"
CONFIG_FILE="$PWD/scripts/swagger/config/cosmos-$TYPE-config.json"
OUTPUT_FILE="$PWD/docs/static/cosmos-$TYPE-openapi.json"

COSMOS_SDK_VERSION="v0.47.8"

rm -rf ./tmp-swagger-gen
mkdir -p ./tmp-swagger-gen

# clone cosmos-sdk repo
git clone -b $COSMOS_SDK_VERSION --depth 1 https://github.com/cosmos/cosmos-sdk.git ./tmp-swagger-gen/cosmos-sdk

cd ./tmp-swagger-gen/cosmos-sdk/proto
pwd
proto_dirs=$(find ./cosmos -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 2 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then
    buf generate --template buf.gen.swagger.yaml $query_file
  fi
done

cd ..
# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine $CONFIG_FILE -o $OUTPUT_FILE --continueOnConflictingPaths true --includeDefinitions true
cd ../..
# clean swagger files
rm -rf ./tmp-swagger-gen