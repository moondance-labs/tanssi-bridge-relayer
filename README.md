# Getting started

## Override configuration

To make sure to use latest contracts, run the following:

```bash
./add_overridden_contracts.sh
```

## Deploy BeeyClient

Go to `snowbridge/contracts`, first rename `.env.example` to `.env` and add the following variables:

```bash
export RPC_URL=
export ETHERSCAN_API_KEY=
export PRIVATE_KEY=
```

To deploy the BeefyClient a `beefy-state.json` file is needed. Then copy it to `snowbridge/contracts` folder.

Load your .env:

```bash
source .env
```

Then run the following command:

```bash
forge script scripts/DeployBeefyLocal.sol --rpc-url ${RPC_URL} --private-key ${PRIVATE_KEY} --slow --skip-simulation --broadcast --verify --etherscan-api-key ${ETHERSCAN_API_KEY}
```

## Deploy Gateway

Go to `snowbridge/contracts`, fill in your `.env` file the following variable:

```bash
export BEEFY_CLIENT_CONTRACT_ADDRESS=
```

Load your .env:

```bash
source .env
```

Then run the following command:

```bash
forge script overridden_contracts/scripts/DeployLocalGateway.sol --rpc-url ${RPC_URL} --private-key ${PRIVATE_KEY} --slow --skip-simulation --verify --etherscan-api-key ${ETHERSCAN_API_KEY}
```

## Deploy Both Gateway and BeefyClient

Go to `snowbridge/contracts`, fill in your `.env` file the following variable:

```bash
export RPC_URL=
export ETHERSCAN_API_KEY=
export PRIVATE_KEY=
```

Load your .env:

```bash
source .env
```

Then run the following command:

```bash
forge script scripts/DeployLocal.sol --rpc-url ${RPC_URL} --private-key ${PRIVATE_KEY} --slow --skip-simulation --broadcast --verify --etherscan-api-key ${ETHERSCAN_API_KEY}
```

