# -----------------------------------------------------------------------------
# This Makefile is used for building your AVS application.
#
# It contains basic targets for building the application, installing dependencies,
# and building a Docker container.
#
# Modify each target as needed to suit your application's requirements.
# -----------------------------------------------------------------------------

GO = $(shell which go)
OUT = ./bin

build: deps
	@mkdir -p $(OUT) || true
	@echo "Building binaries..."
	go build -o $(OUT)/performer ./avs/cmd/main.go
	cd contracts && forge build

deps:
	GOPRIVATE=github.com/Layr-Labs/* go mod tidy
	cd contracts && forge install

build/container:
	./.hourglass/scripts/buildContainer.sh

anvil:
	cd contracts && anvil --fork-url $(MAINNET_RPC_URL) --fork-block-number 22396947

deploy:
	cd contracts && make deploy-task-mailbox RPC_URL="127.0.0.1:8545"
	cd contracts && make deploy-avs-l1-contracts AVS_ADDRESS='0x70997970C51812dc3A010C7d01b50e0d17dc79C8' RPC_URL="127.0.0.1:8545"
	cd contracts && make setup-avs-l1 TASK_AVS_REGISTRAR_ADDRESS='0xf4c5C29b14f0237131F7510A51684c8191f98E06' RPC_URL="127.0.0.1:8545"
	cd contracts && make deploy-avs-l2-contracts RPC_URL="127.0.0.1:8545"
	cd contracts && make setup-avs-task-mailbox-config TASK_MAILBOX_ADDRESS='0x7306a649B451AE08781108445425Bd4E8AcF1E00' CERTIFICATE_VERIFIER_ADDRESS='0xc91B651f770ed996a223a16dA9CCD6f7Df56C987' TASK_HOOK_ADDRESS='0x934A389CaBFB84cdB3f0260B2a4FD575b8B345A3' RPC_URL="127.0.0.1:8545"

run:
	cd contracts && make create-task TASK_MAILBOX_ADDRESS='0x7306a649B451AE08781108445425Bd4E8AcF1E00' AVS_ADDRESS='0x70997970C51812dc3A010C7d01b50e0d17dc79C8' RPC_URL="127.0.0.1:8545"
