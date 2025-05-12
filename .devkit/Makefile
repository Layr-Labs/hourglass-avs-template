# This Makefile is specifically used as an interface for the devkit-cli.
#
# If you need to modify how your project is build, you should modify the main Makefile

.PHONY: build
build:
	BUILD_CONTAINER=true ./.hourglass/scripts/build.sh
	cd contracts && forge build

.PHONY: anvil
anvil:
	cd contracts && anvil --fork-url $(FORK_URL) --chain-id 31337 --block-time 12

.PHONY: deploy
deploy:
	cd contracts && PRIVATE_KEY_DEPLOYER="$(PRIVATE_KEY_DEPLOYER)" make deploy-task-mailbox RPC_URL="$(L2_RPC_URL)" ENVIRONMENT="$(ENVIRONMENT)"
	cd contracts && PRIVATE_KEY_DEPLOYER="$(PRIVATE_KEY_DEPLOYER)" make deploy-avs-l1-contracts RPC_URL="$(L1_RPC_URL)" ENVIRONMENT="$(ENVIRONMENT)" AVS_ADDRESS='$(AVS_ADDRESS)' ALLOCATION_MANAGER_ADDRESS='$(ALLOCATION_MANAGER_ADDRESS)'
	cd contracts && PRIVATE_KEY_AVS="$(PRIVATE_KEY_AVS)" make setup-avs-l1 RPC_URL="$(L1_RPC_URL)" ENVIRONMENT="$(ENVIRONMENT)" ALLOCATION_MANAGER_ADDRESS='$(ALLOCATION_MANAGER_ADDRESS)' METADATA_URI="$(METADATA_URI)" AGGREGATOR_OPERATOR_SET_ID=$(AGGREGATOR_OPERATOR_SET_ID) AGGREGATOR_STRATEGIES='$(AGGREGATOR_STRATEGIES)' EXECUTOR_OPERATOR_SET_ID=$(EXECUTOR_OPERATOR_SET_ID) EXECUTOR_STRATEGIES='$(EXECUTOR_STRATEGIES)'
	cd contracts && PRIVATE_KEY_DEPLOYER="$(PRIVATE_KEY_DEPLOYER)" make deploy-avs-l2-contracts RPC_URL="$(L2_RPC_URL)" ENVIRONMENT="$(ENVIRONMENT)"
	cd contracts && PRIVATE_KEY_AVS="$(PRIVATE_KEY_AVS)" make setup-avs-task-mailbox-config RPC_URL="$(L2_RPC_URL)" ENVIRONMENT="$(ENVIRONMENT)" AGGREGATOR_OPERATOR_SET_ID=$(AGGREGATOR_OPERATOR_SET_ID) EXECUTOR_OPERATOR_SET_ID=$(EXECUTOR_OPERATOR_SET_ID) TASK_SLA=$(TASK_SLA)
	./.hourglass/scripts/run.sh

.PHONY: run
run:
	cd contracts && PRIVATE_KEY_APP="$(PRIVATE_KEY_APP)" make create-task RPC_URL="$(L2_RPC_URL)" ENVIRONMENT="$(ENVIRONMENT)" AVS_ADDRESS='$(AVS_ADDRESS)' EXECUTOR_OPERATOR_SET_ID=$(EXECUTOR_OPERATOR_SET_ID) PAYLOAD='$(PAYLOAD)'
