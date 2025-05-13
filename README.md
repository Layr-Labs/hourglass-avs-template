# hourglass-avs-template

This template is for building an AVS with the EigenLayer Hourglass framework. It provides a basic structure and configuration files to get you started.

It is also compatibile with the devkit-cli which will help you build, test, and deploy your AVS.

## Basic Structure

This template includes a basic Go program that uses the Hourglass framework to get you started along with some sample configs, a Dockerfile and basic build scripts.

```bash
.
|-- .gitignore
|-- .gitmodules
|-- .devkit
|   |-- Makefile
|   |-- README.md
|-- .hourglass
|   |-- build.yaml
|   |-- config
|   |   |-- aggregator.yaml
|   |   `-- executor.yaml
|   `-- scripts
|       |-- build.sh
|       |-- buildContainer.sh
|       |-- init.sh
|       `-- run.sh
|-- Dockerfile
|-- Makefile
|-- README.md
|-- avs
|   `-- cmd
|       `-- main.go
|-- bin
|   `-- performer
|-- config
|   `-- README.md
|-- contracts
|   |-- lib
|   |-- script
|   |   `-- local
|   |       |-- deploy
|   |       |   |-- DeployAVSL1Contracts.s.sol
|   |       |   |-- DeployAVSL2Contracts.s.sol
|   |       |   |-- DeployTaskMailbox.s.sol
|   |       |   `-- DeployTaskMailboxConfig.s.sol
|   |       |-- output
|   |       |   |-- deploy_avs_l1_output.json
|   |       |   |-- deploy_avs_l2_output.json
|   |       |   |-- deploy_hourglass_output.json
|   |       |-- run
|   |       |   |-- CreateTask.s.sol
|   |       |-- setup
|   |       |   |-- SetupAVSL1.s.sol
|   |       |   `-- SetupAVSTaskMailboxConfig.s.sol
|   |-- src
|   |   |-- l1-contracts
|   |   |   |-- TaskAVSRegistrar.sol
|   |   |-- l2-contracts
|   |   |   |-- AVSTaskHook.sol
|   |   |   `-- BN254CertificateVerifier.sol
|   |-- test
|   |   `-- TaskAVSRegistrar.t.sol
|   |-- .env.example
|   |-- .gitignore
|   |-- .gitmodules
|   |-- foundry.toml
|   |-- Makefile
|   `-- README.md
|-- go.mod
`-- go.sum

```

## Getting Started

Follow these steps to set up and run your AVS with the Hourglass framework:

### 0. Prerequisites

Follow the instructions in the [devkit-cli](https://github.com/Layr-Labs/devkit-cli) README to install the devkit.

### 1. AVS Logic

Update the `avs/cmd/main.go` file to implement your offchain AVS logic. Additionally, update the contracts in `contracts/src` to add any additional logic (only if needed).

### 2. Build the AVS project

Build the project and the contracts:

```bash
devkit avs build
```

### 3. Update the AVS config

Read or modify eigen.yaml configuration

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
devkit avs run PAYLOAD='<PAYLOAD>'
```

