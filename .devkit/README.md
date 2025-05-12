# DevKit Commands

This README explains how to use the commands in the `.devkit/Makefile`. These commands provide a simplified interface for running the AVS deployment and task creation scripts.

## Prerequisites

- [Foundry](https://book.getfoundry.sh/) installed
- [Anvil](https://book.getfoundry.sh/anvil/) (comes with Foundry) 
- `make` installed
- Access to an Ethereum RPC URL (for forking mainnet)

## Available Commands

### 1. Build the Project

Build the contracts and dependencies:

```sh
make -f .devkit/Makefile build
```

### 2. Run a Local Ethereum Node with Anvil

Start a local Ethereum node forked from mainnet:

```sh
make -f .devkit/Makefile anvil FORK_URL="<YOUR_ETHEREUM_RPC_URL>"
```

This will run a local node at `127.0.0.1:8545` with chain ID 31337 and a 12-second block time.

### 3. Deploy the TaskMailbox Contract

Deploy the TaskMailbox contract:

```sh
make -f .devkit/Makefile deploy-task-mailbox \
  L2_RPC_URL="127.0.0.1:8545" \
  PRIVATE_KEY_DEPLOYER="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" \
  ENVIRONMENT="local"
```

This uses predefined private keys from Anvil for deployment. Make sure your Anvil node is running before executing this command.

### 4. Deploy and set up the Contracts

The deploy command executes all the deployment and setup steps in sequence:

1. Deploys AVS L1 contracts
2. Sets up AVS on L1
3. Deploys AVS L2 contracts
4. Sets up AVS Task Mailbox configuration

Run it with:

```sh
make -f .devkit/Makefile deploy \
  L1_RPC_URL="127.0.0.1:8545" \
  L2_RPC_URL="127.0.0.1:8545" \
  PRIVATE_KEY_DEPLOYER="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" \
  PRIVATE_KEY_AVS="0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d" \
  ENVIRONMENT="local" \
  AVS_ADDRESS='0x70997970C51812dc3A010C7d01b50e0d17dc79C8' \
  ALLOCATION_MANAGER_ADDRESS='0x948a420b8CC1d6BFd0B6087C2E7c344a2CD0bc39' \
  METADATA_URI="TestAVS" \
  AGGREGATOR_OPERATOR_SET_ID=0 \
  AGGREGATOR_STRATEGIES='["0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7","0x93c4b944D05dfe6df7645A86cd2206016c51564D"]' \
  EXECUTOR_OPERATOR_SET_ID=1 \
  EXECUTOR_STRATEGIES='["0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7","0x93c4b944D05dfe6df7645A86cd2206016c51564D"]' \
  TASK_SLA=60
```

This uses predefined private keys from Anvil for deployment. Make sure your Anvil node is running before executing this command.

### 5. Run the AVS

Run the Aggregator and Executor in docker containers:

```sh
make -f .devkit/Makefile run-avs
```

### 6. Create a Task

Create a task on the TaskMailbox contract:

```sh
make -f .devkit/Makefile run \
  L2_RPC_URL="127.0.0.1:8545" \
  PRIVATE_KEY_APP="0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a" \
  ENVIRONMENT="local" \
  AVS_ADDRESS='0x70997970C51812dc3A010C7d01b50e0d17dc79C8' \
  EXECUTOR_OPERATOR_SET_ID=1 \
  PAYLOAD='0x0000000000000000000000000000000000000000000000000000000000000005'
```

This will create a task on the TaskMailbox contract with the specified AVS address, operator set ID, and payload.

## Developer Notes

The commands in this Makefile are set up for local development and testing. They use:

- Hardcoded private keys from Anvil for simplicity
- Predefined addresses for contracts
- Default values for parameters

For production deployment, you would need to:
1. Set up appropriate private keys
2. Use your own contract addresses
3. Customize parameters according to your needs

Refer to `contracts/script/local/README.md` for more detailed information about each script and parameter. 