#!/usr/bin/env bash
# possible values: "base" or "tx"
TYPE=${1:-tx}

echo "Generating Cosmos openapi for '$TYPE'"
CONFIG_FILE="$PWD/scripts/swagger/config/cosmos-$TYPE-config.json"
OUTPUT_FILE="$PWD/docs/static/cosmos-$TYPE-openapi.json"

COSMOS_SDK_VERSION="v0.44.4"

rm -rf ./tmp-swagger-gen
mkdir -p ./tmp-swagger-gen

# clone cosmos-sdk repo
git clone -b $COSMOS_SDK_VERSION --depth 1 https://github.com/cosmos/cosmos-sdk.git ./tmp-swagger-gen/cosmos-sdk

# generate openapi
pushd ./tmp-swagger-gen/cosmos-sdk
../../scripts/swagger/protoc-swagger-gen.sh $CONFIG_FILE $OUTPUT_FILE
popd

rm -rf ./tmp-swagger-gen