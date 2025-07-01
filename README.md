# Hourglass AVS Template

## Getting Started

After generating this project with `devkit avs create`, follow these steps to build and run your AVS:

```bash
# Build your AVS Go code and contracts
devkit avs build

# Start the devnet (this will deploy contracts and start the Hourglass infrastructure). 
# `--skip-avs-run` will skip running your AVS Performer locally, allowing you to run it separately.
devkit avs devnet start [--skip-avs-run]

# If you ran devnet start with `--skip-avs-run`, you can now run your AVS Performer separately:
devkit avs run
```

## Where Your Code Goes

### Go Code - `cmd/main.go`

This is your AVS Performer implementation - the core business logic that processes tasks. The starter file includes the bare minimum needed to get you started:

- `ValidateTask()` - Validate incoming task requests
- `HandleTask()` - Process tasks and return results

This is just a starting structure. Feel free to restructure the code however you see fit for your AVS requirements.

### Smart Contracts - `contracts/src/`

Your custom contracts go here. The template includes:

- `HelloWorld.sol` - An example contract you can delete if not needed
- `l1-contracts/TaskAVSRegistrar.sol` - L1 operator registration (extend as needed)
- `l2-contracts/AVSTaskHook.sol` - Task lifecycle validation (extend as needed)

#### Deploying Your Contracts

Wire up your contracts in `contracts/script/DeployMyContracts.s.sol`. This script is automatically called during `devkit avs devnet start`:

```solidity
// Deploy your contract
CustomContract customContract = new CustomContract();

// Add it to the output so devkit can track it
Output[] memory outputs = new Output[](1);
outputs[0] = Output({name: "CustomContract", contractAddress: address(customContract)});
_writeOutputToJson(environment, outputs);
```

## What is Hourglass?

Hourglass is a framework for building task-based EigenLayer AVSs. It provides a batteries-included experience with onchain components (TaskMailbox, TaskAVSRegistrar, AVSTaskHook) and offchain components (Aggregator, Executor, Performer) that work together to handle task distribution, execution, and result aggregation.

For more details, visit the framework repository: https://github.com/Layr-Labs/hourglass-monorepo

## ⚠️ Warning: This is Alpha, non-audited code ⚠️
Hourglass is in active development and is not yet audited. Use at your own risk.
