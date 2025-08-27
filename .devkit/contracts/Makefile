# Build the project
.PHONY: build
build:
	forge clean
	forge build

# Test the project
.PHONY: test
test:
	forge test -vvv

# Deploy AVS L1 Contracts
.PHONY: deploy-avs-l1-contracts
deploy-avs-l1-contracts:
	forge script script/deploy/DeployAVSL1Contracts.s.sol \
		--rpc-url $(RPC_URL) \
		--broadcast \
		--sig "run(string, address, address, address, address, uint32, uint32, address[])" $(ENVIRONMENT) $(AVS_ADDRESS) $(ALLOCATION_MANAGER_ADDRESS) $(KEY_REGISTRAR_ADDRESS) $(PERMISSION_CONTROLLER_ADDRESS) $(AGGREGATOR_OPERATOR_SET_ID) $(EXECUTOR_OPERATOR_SET_ID) $(AGGREGATOR_OPERATORS) \
		--slow \
		-vvvv

# Deploy AVS L2 Contracts
.PHONY: deploy-avs-l2-contracts
deploy-avs-l2-contracts:
	forge script script/deploy/DeployAVSL2Contracts.s.sol \
		--rpc-url $(RPC_URL) \
		--broadcast \
		--sig "run(string)" $(ENVIRONMENT) \
		--slow \
		-vvvv

# Setup AVS Task Mailbox Config
.PHONY: setup-avs-task-mailbox-config
setup-avs-task-mailbox-config:
	forge script script/setup/SetupAVSTaskMailboxConfig.s.sol \
		--rpc-url $(RPC_URL) \
		--broadcast \
		--sig "run(string, uint32, uint96, uint8, string)" $(ENVIRONMENT) $(EXECUTOR_OPERATOR_SET_ID) $(TASK_SLA) $(CURVE_TYPE) '$(CONTEXT)' \
		--slow \
		-vvvv

# Deploy Custom L1 Contracts
.PHONY: deploy-custom-contracts-l1
deploy-custom-contracts-l1:
	forge script $(shell pwd)/../../contracts/script/DeployMyL1Contracts.s.sol \
		--lib-paths . \
		--rpc-url $(RPC_URL) \
		--broadcast \
		--sig "run(string, string)" "$(ENVIRONMENT)" '$(CONTEXT)'\
		--slow \
		-vvvv

# Deploy Custom L2 Contracts
.PHONY: deploy-custom-contracts-l2
deploy-custom-contracts-l2:
	forge script $(shell pwd)/../../contracts/script/DeployMyL2Contracts.s.sol \
		--lib-paths . \
		--rpc-url $(RPC_URL) \
		--broadcast \
		--sig "run(string, string)" "$(ENVIRONMENT)" '$(CONTEXT)'\
		--slow \
		-vvvv

# Create Task
.PHONY: create-task
create-task:
	forge script script/call/CreateTask.s.sol \
		--rpc-url $(RPC_URL) \
		--broadcast \
		--sig "run(string, address, address, uint32, bytes)" $(ENVIRONMENT) $(TASK_MAILBOX_ADDRESS) $(AVS_ADDRESS) $(EXECUTOR_OPERATOR_SET_ID) $(PAYLOAD) \
		--slow \
		-vvvv

# Helper message
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build - Build the project"
	@echo "  make test - Test the project"
	@echo "  make deploy-avs-l1-contracts AVS_ADDRESS=0x... - Deploy AVS L1 Contracts"
	@echo "  make deploy-avs-l2-contracts - Deploy AVS L2 Contracts"
	@echo "  make setup-avs-task-mailbox-config TASK_MAILBOX_ADDRESS=0x... CERTIFICATE_VERIFIER_ADDRESS=0x... TASK_HOOK_ADDRESS=0x... - Setup AVS Task Mailbox Config"
	@echo "  make create-task TASK_MAILBOX_ADDRESS=0x... AVS_ADDRESS=0x... VALUE=5 - Create Task"
	@echo ""
	@echo "Note: Make sure to set RPC_URL and PRIVATE_KEY in your environment or .env file" 
