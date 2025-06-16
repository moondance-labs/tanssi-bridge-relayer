#!/bin/bash

# Exit on any error
set -e

echo -e "\033[96mUpdating contract interfaces from snowbridge fork, make sure you are on the correct branch...\033[0m"

# Assuming that we have a folder with cloned snowbridge fork repo: moondance-labs/snowbridge
FOLDER="../snowbridge/contracts"
if [ ! -d $FOLDER ]; then
  echo "Please clone the snowbridge fork repo to the parent directory as 'snowbridge'."
  exit 1
fi
(cd $FOLDER && forge build)

jq .abi "$FOLDER/out/BeefyClient.sol/BeefyClient.json" | abigen --abi - --type BeefyClient --pkg contracts --out relays/contracts/beefy_client.go
jq .abi "$FOLDER/out/IGateway.sol/IGateway.json" | abigen --abi - --type Gateway --pkg contracts --out relays/contracts/gateway.go