#!/usr/bin/env bash

echo "Generating DCL openapi ..."
CONFIG_FILE="$PWD/scripts/swagger/config/dcl-config.json"
OUTPUT_FILE="$PWD/docs/static/openapi.yml"

./scripts/swagger/protoc-swagger-gen.sh $CONFIG_FILE $OUTPUT_FILE