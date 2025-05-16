# hourglass-avs-template

## What is Hourglass?

Hourglass is a framework for building a task-based EigenLayer AVS, providing AVS developers a batteries-included experience to get started quickly. It includes a set of tools and libraries that simplify the process of building, deploying, and managing AVS projects.

![](docs/images/hourglass-architecture_v.01.0.svg)

Hourglass as a framework has onchain and offchain components that work together to enable a task-based AVS.

### Onchain Components

#### TaskMailbox

The TaskMailbox is a singleton eigenlayer hourglass contract on L1 or L2 that is responsible for:

* Allowing users/apps to create tasks.
* Managing the lifecycle of tasks.
* Verifying the results of tasks and making it available for users/apps to query.
* Allowing AVSs to manage their TaskMailbox configurations.

#### TaskAVSRegistrar

The TaskAVSRegistrar is an instanced (per-AVS) eigenlayer middleware contract on L1 that is responsible for:

* Handling operator registration for specific operator sets of your AVS.
* Providing the offchain components with BLS public keys and socket endpoints for the Aggregator and Executor operators.

It works by default, but can be extended to include additional onchain logic for your AVS.

#### AVSTaskHook

The AVSTaskHook is an instanced (per-AVS) eigenlayer hourglass contract on L1 or L2 that is responsible for:

* Validating the task lifecycle.
* Creating fee markets for your AVS.

It's empty by default and works out of the box, but can be extended to include additional onchain validation logic for your AVS.

#### CertificateVerifier

The CertificateVerifier is an instanced (per-AVS) eigenlayer middleware contract on L1 or L2 that is responsible for:

* Verifying the validity of operator certificates.
* Verifying stake threshold requirements for operator sets.

### Offchain Components

#### Aggregator

The Aggregator is responsible for:

* Listening to events from the Mailbox contract on chain for new tasks
* Discovering Executors by querying the AVSRegistrar contract (via the EigenLayer Allocation manager), retrieving their metadata containing a BLS public key and a "socket" (url) endpoint that references the Executor's gRPC server.
* Distributing tasks to Executors by sending a gRPC request to the Executor's socket endpoint, including the task payload and a signature of the payload signed by the Aggregator. This is so the Executor can validate the message is coming from the expected Aggregator.
* Aggregates results from Executors until a signing threshold has been met
* Publish the result back to the Mailbox contract

#### Executor

The Executor is responsible for:
* Launching and managing Performer containers that execute the tasks
* Listening to gRPC requests from the Aggregator for new tasks
* Forwarding the task to the correct Performer
* Signing the result of the task with its BLS private key and sending it back to the Aggregator


#### Performer

The Performer is the component the AVS is responsible for building. At a high level, it is a simple gRPC server that listens for tasks, runs them and returns the results to the Executor.

The Hourglass framework provides all of the boilerplate and server code for your Performer; you simply need to fill in the logic to handle tasks for your AVS!

## What does this template give me?

This template provides a basic structure for building an AVS with the Hourglass framework. It includes:

* A stub of Go code for your Performer to get you started. Simply fill in the commented out areas with your own AVS logic
* Default `TaskAVSRegistrar` and `AVSTaskHook` avs contracts that work out of the box. Simply extend them if you need to add additional onchain logic.
* All the dependent contracts for the framework to work and scripts to deploy them. The scripts will be managed by the Devkit CLI.
* A docker-compose stack to run an Aggregator and Executor locally to test your AVS. Both the Aggregator and Executor will be run by EigenLayer Operators when you launch your AVS, so we've given you the full stack to run locally to make development and testing easier.
* Hooks that integrate with the Devkit CLI. The Devkit CLI is a command line tool that will make your development experience faster and easier by automating common tasks like building, deploying, and running your AVS.


## Basic Structure

This template includes a basic Go program and smart contracts that uses the Hourglass framework to get you started along with some default configs.

```bash
.
|-- .gitignore
|-- .gitmodules
|-- .devkit
|   |-- scripts
|       |-- build
|       |-- call
|       |-- deployContracts
|       |-- getOperatorRegistrationMetadata
|       |-- getOperatorSets
|       |-- init
|       |-- run
|-- .hourglass
|   |-- build.yaml
|   |-- docker-compose.yml
|   |-- context
|   |   |-- devnet.yaml
|   |-- config
|   |   |-- aggregator.yaml
|   |   |-- executor.yaml
|   |-- scripts
|       |-- build.sh
|       |-- buildContainer.sh
|       |-- init.sh
|       |-- run.sh
|-- Dockerfile
|-- Makefile
|-- README.md
|-- avs
|   |-- cmd
|       |-- main.go
|-- contracts
|   |-- lib
|   |-- script
|   |   |-- devnet
|   |       |-- deploy
|   |       |   |-- DeployAVSL1Contracts.s.sol
|   |       |   |-- DeployAVSL2Contracts.s.sol
|   |       |   |-- DeployTaskMailbox.s.sol
|   |       |-- output
|   |       |   |-- deploy_avs_l1_output.json
|   |       |   |-- deploy_avs_l2_output.json
|   |       |   |-- deploy_hourglass_core_output.json
|   |       |-- run
|   |       |   |-- CreateTask.s.sol
|   |       |-- setup
|   |       |   |-- SetupAVSL1.s.sol
|   |       |   |-- SetupAVSTaskMailboxConfig.s.sol
|   |-- src
|   |   |-- l1-contracts
|   |   |   |-- TaskAVSRegistrar.sol
|   |   |-- l2-contracts
|   |   |   |-- AVSTaskHook.sol
|   |   |   |-- BN254CertificateVerifier.sol
|   |-- test
|   |   |-- TaskAVSRegistrar.t.sol
|   |-- foundry.toml
|   |-- Makefile
|-- go.mod
|-- go.sum
```

## Getting Started

Follow these steps to set up and run your AVS with the Hourglass framework:

### 0. Prerequisites

Follow the instructions in the [devkit-cli](https://github.com/Layr-Labs/devkit-cli) README to install the devkit.

### 1. AVS Logic

Update the `cmd/main.go` file to implement your offchain AVS logic. 
Update the `TaskAVSRegistrar.sol` and `AVSTaskHook.sol` contracts in `contracts/src` to add any additional onchain logic (only if needed).

### 2. Build the AVS project

Build the project and the contracts:

```bash
devkit avs build
```

### 3. Update the AVS config

Read or modify the AVS configuration to set the correct parameters for your AVS.

```bash
devkit avs config
```

### 4. Start the AVS devnet

Start the AVS devnet, deploy and set up the contracts, and run the AVS:

```bash
devkit avs devnet start
```

### 5. Create a task

Create a task on the TaskMailbox contract:

```bash
devkit avs call --params payload="<payload>"
```

