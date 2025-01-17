#!/bin/bash

# Exit on any error
set -e

# Copy overridden contracts into main contracts directory
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    cp -R ./overridden_contracts/src/ ./snowbridge/contracts/src
    cp -R ./overridden_contracts/test/ ./snowbridge/contracts/test
else
    # Linux
    cp -TR ./overridden_contracts/src/ ./snowbridge/contracts/src
    cp -TR ./overridden_contracts/test/ ./snowbridge/contracts/test
fi
# Compile the resulting contracts
pushd ./snowbridge/contracts
forge build
popd
