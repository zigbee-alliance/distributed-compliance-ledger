#!/bin/bash

set -euox pipefail

echo "Generating gogo proto code"
cd proto

buf generate --template buf.gen.gogo.yaml

cd ..

# move proto files to the right places
cp -r github.com/zigbee-alliance/distributed-compliance-ledger/* ./
rm -rf github.com
