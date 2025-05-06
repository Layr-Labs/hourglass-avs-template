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

## Contracts

The `/contracts` directory contains the Hourglass contracts template as a git submodule. This provides the smart contracts needed for your AVS to interact with EigenLayer. The contracts are maintained in a separate repository at [github.com/Layr-Labs/hourglass-contracts-template](https://github.com/Layr-Labs/hourglass-contracts-template).

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
