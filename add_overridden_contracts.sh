#!/bin/bash

# Exit on any error
set -e

# Copy overridden contracts into main contracts directory
cp  -TR ./overridden_contracts/src/ ./snowbridge/contracts/src
cp  -TR ./overridden_contracts/test/ ./snowbridge/contracts/test

# Compile the resulting contracts
pushd ./snowbridge/contracts
forge build
popd
