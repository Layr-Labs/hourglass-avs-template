# DevKit Commands

This README explains how to use the commands in the `.devkit/Makefile`. These commands provide a simplified interface for running the AVS deployment and task creation scripts.

## Prerequisites

- [Foundry](https://book.getfoundry.sh/) installed
- [Anvil](https://book.getfoundry.sh/anvil/) (comes with Foundry) 
- `make` installed
- Access to an Ethereum RPC URL (for forking mainnet)

## Available Commands

### Building the Project

Build the contracts and dependencies:

```sh
make -f .devkit/Makefile build
```

### Running a Local Ethereum Node with Anvil

Start a local Ethereum node forked from mainnet:

```sh
make -f .devkit/Makefile anvil RPC_URL="<YOUR_ETHEREUM_RPC_URL>"
```

This will run a local node at `127.0.0.1:8545` with chain ID 31337 and a 12-second block time.

### Deploying the Contracts

The deploy command executes all the deployment steps in sequence:

1. Deploys the TaskMailbox contract
2. Deploys AVS L1 contracts
3. Sets up AVS on L1
4. Deploys AVS L2 contracts
5. Sets up AVS Task Mailbox configuration

Run it with:

```sh
make -f .devkit/Makefile deploy
```

This uses predefined private keys from Anvil for deployment. Make sure your Anvil node is running before executing this command.

### Creating a Task

After deployment, you can create a task with:

```sh
make -f .devkit/Makefile run
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