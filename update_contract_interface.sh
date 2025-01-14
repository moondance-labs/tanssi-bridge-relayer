#!/bin/bash

# Exit on any error
set -e

# Copy overridden contracts into main contracts directory
cp  -R ./overridden_contracts/src/ ./snowbridge/contracts/src
cp  -R ./overridden_contracts/test/ ./snowbridge/contracts/test

# Compile the resulting contracts
pushd ./snowbridge/contracts
forge build
popd

# Create Go interface to it in main directory
jq .abi ./snowbridge/contracts/out/BeefyClient.sol/BeefyClient.json | abigen --abi - --type BeefyClient --pkg contracts --out ./relays/contracts/beefy_client.go
jq .abi ./snowbridge/contracts/out/IGateway.sol/IGateway.json | abigen --abi - --type Gateway --pkg contracts --out ./relays/contracts/gateway.go
