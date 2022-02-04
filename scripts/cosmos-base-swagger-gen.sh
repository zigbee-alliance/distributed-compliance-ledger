#!/usr/bin/env bash

# parts are copied from https://github.com/cosmos/cosmos-sdk/blob/master/scripts/protoc-swagger-gen.sh

# possible values: "base" or "tx"
TYPE=${1:-base}

echo "Generating Cosmos openapi for '$TYPE'"

COSMOS_SDK_VERSION="v0.44.4"

set -eo pipefail

rm -rf ./tmp-swagger-gen
mkdir -p ./tmp-swagger-gen

# generate for cosmos-sdk `base` or `tx` proto only
git clone -b $COSMOS_SDK_VERSION --depth 1 git@github.com:cosmos/cosmos-sdk.git ./tmp-swagger-gen/cosmos-sdk
cd ./tmp-swagger-gen/cosmos-sdk
proto_dirs=$(find proto/cosmos/$TYPE -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
out_file="../../docs/static"

for dir in $proto_dirs; do

  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then
    protoc  \
      -I "proto" \
      -I "third_party/proto" \
      "$query_file" \
      --openapiv2_out "$out_file" \
      --openapiv2_opt logtostderr=true \
      --openapiv2_opt allow_merge=true \
      --openapiv2_opt fqn_for_openapi_name=true \
      --openapiv2_opt simple_operation_ids=true \
      --openapiv2_opt Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:.
  fi
done

DEST_FILE="cosmos-$TYPE-openapi.json"
DESC="A REST interface for $TYPE service endpoints"

mv "$out_file/apidocs.swagger.json" "$out_file/$DEST_FILE"
sed -i "s/\"version\": \"version not set\"/\"version\": \"$COSMOS_SDK_VERSION\",\"description\": \"$DESC\"/g" "$out_file/$DEST_FILE"

# clean swagger files
cd ../..
rm -rf ./tmp-swagger-gen
