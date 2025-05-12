# This Makefile is specifically used as an interface for the devkit-cli.
#
# If you need to modify how your project is build, you should modify the main Makefile

.PHONY: build
build:
	BUILD_CONTAINER=true ../.hourglass/scripts/build.sh
	cd contracts && forge build

.PHONY: anvil
anvil:
	cd contracts && anvil --fork-url $(RPC_URL) --chain-id 31337 --block-time 12

.PHONY: deploy
deploy:
	cd contracts && PRIVATE_KEY_DEPLOYER="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" make deploy-task-mailbox RPC_URL="127.0.0.1:8545" ENVIRONMENT="local"
	cd contracts && PRIVATE_KEY_DEPLOYER="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" make deploy-avs-l1-contracts RPC_URL="127.0.0.1:8545" ENVIRONMENT="local" AVS_ADDRESS='0x70997970C51812dc3A010C7d01b50e0d17dc79C8' ALLOCATION_MANAGER_ADDRESS='0x948a420b8CC1d6BFd0B6087C2E7c344a2CD0bc39'
	cd contracts && PRIVATE_KEY_AVS="0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d" make setup-avs-l1 RPC_URL="127.0.0.1:8545" ENVIRONMENT="local" ALLOCATION_MANAGER_ADDRESS='0x948a420b8CC1d6BFd0B6087C2E7c344a2CD0bc39' METADATA_URI="TestAVS" AGGREGATOR_OPERATOR_SET_ID=0 AGGREGATOR_STRATEGIES='["0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7","0x93c4b944D05dfe6df7645A86cd2206016c51564D"]' EXECUTOR_OPERATOR_SET_ID=1 EXECUTOR_STRATEGIES='["0xaCB55C530Acdb2849e6d4f36992Cd8c9D50ED8F7","0x93c4b944D05dfe6df7645A86cd2206016c51564D"]'
	cd contracts && PRIVATE_KEY_DEPLOYER="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" make deploy-avs-l2-contracts RPC_URL="127.0.0.1:8545" ENVIRONMENT="local"
	cd contracts && PRIVATE_KEY_AVS="0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d" make setup-avs-task-mailbox-config RPC_URL="127.0.0.1:8545" ENVIRONMENT="local" AGGREGATOR_OPERATOR_SET_ID=0 EXECUTOR_OPERATOR_SET_ID=1 TASK_SLA=60
	../.hourglass/scripts/run.sh

.PHONY: run
run:
	cd contracts && PRIVATE_KEY_APP="0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a" make create-task RPC_URL="127.0.0.1:8545" ENVIRONMENT="local" AVS_ADDRESS='0x70997970C51812dc3A010C7d01b50e0d17dc79C8' EXECUTOR_OPERATOR_SET_ID=1 PAYLOAD='0x0000000000000000000000000000000000000000000000000000000000000005'
