# hourglass-avs-template

This template is for building an AVS with the EigenLayer Hourglass framework. It provides a basic structure and configuration files to get you started.

It is also compatibile with the devkit-cli which will help you build, test, and deploy your AVS.

## Basic Structure

This template includes a basic Go program that uses the Hourglass framework to get you started along with some sample configs, a Dockerfile and basic build scripts.

```bash
.
|-- .gitignore
|-- .gitmodules
|-- .hourglass
|   |-- build.yaml
|   |-- config
|   |   |-- aggregator.yaml
|   |   `-- executor.yaml
|   `-- scripts
|       |-- build.sh
|       |-- init.sh
|       `-- run.sh
|-- Dockerfile
|-- Makefile
|-- Makefile.Devkit
|-- README.md
|-- avs
|   `-- cmd
|       `-- main.go
|-- bin
|   `-- performer
|-- config
|   `-- README.md
|-- contracts
|-- go.mod
`-- go.sum

```

### Contracts

The `/contracts` directory contains the Hourglass contracts template as a git submodule. This provides the smart contracts needed for your AVS to interact with EigenLayer. The contracts are maintained in a separate repository at [github.com/Layr-Labs/hourglass-contracts-template](https://github.com/Layr-Labs/hourglass-contracts-template).


## Getting Started

Follow these steps to set up and run your AVS with the Hourglass framework:

### 1. Install Dependencies

Install all required Go and Foundry dependencies:

```bash
make deps
```

### 2. Build the Application

Compile the application binaries and contracts:

```bash
make build
```

This command creates the executable in the `./bin` directory.

### 3. Local Development Environment

Start a local forked Ethereum node with Anvil (requires a mainnet RPC URL):

```bash
make anvil MAINNET_RPC_URL="<MAINNET_RPC_URL>"
```

Keep this running in a separate terminal window.

### 4. Deploy Contracts

Deploy all necessary contracts to your local development environment:

```bash
make deploy
```

This command:
- Deploys the task mailbox
- Deploys AVS L1 contracts with the specified AVS address
- Sets up the AVS L1 with the EigenLayer core protocol
- Deploys AVS L2 contracts
- Configures the task mailbox with appropriate AVS addresses and configs.

### 5. Run a Task

Create a task in the deployed contracts:

```bash
make run
```

## Interacting with the aggregator

```bash
curl -H 'content-type: application/json' -XPOST localhost:8081/events -d '{ 
        "taskId": "0xtask1",
        "avsAddress": "0xavs1...",
        "operatorSetId": 1,
        "callbackAddr": "0xcallmemaybe",
        "deadline": 300,
        "stakeRequired": 100,
        "payload": "eyAibnVtYmVyVG9CZVNxdWFyZWQiOiA0IH0=",
        "chainId": 1,
        "blockNumber": 12345,
        "blockHash": "0xblockHash"
}'
```
