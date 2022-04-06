#!/usr/bin/env bash
# Copyright 2022
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

TYPE=${1:-tx}
valid_values=(base tx)
if ! printf '%s\0' "${valid_values[@]}" | grep -qxFe "$TYPE"; then
    printf '%s\n' "$TYPE is an invalid value"
    exit 1
fi

echo "Generating Cosmos openapi for '$TYPE'"
CONFIG_FILE="$PWD/scripts/swagger/config/cosmos-$TYPE-config.json"
OUTPUT_FILE="$PWD/docs/static/cosmos-$TYPE-openapi.json"
COSMOS_SDK_VERSION="v0.44.4"

mkdir -p ./tmp-swagger-gen
if [ ! -d ./tmp-swagger-gen/cosmos-sdk ]; then
    git clone -b $COSMOS_SDK_VERSION --depth 1 https://github.com/cosmos/cosmos-sdk.git ./tmp-swagger-gen/cosmos-sdk
fi
trap "rm -rf ./tmp-swagger-gen" EXIT

# generate openapi
pushd ./tmp-swagger-gen/cosmos-sdk || exit
../../scripts/swagger/protoc-swagger-gen.sh "$CONFIG_FILE" "$OUTPUT_FILE"
popd || exit
