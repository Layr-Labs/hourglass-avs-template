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
GIT_HOURGLASS_REPO=git@github.com:Layr-Labs/hourglass-monorepo.git
HOURGLASS_DIR=./temp_external/hourglass-monorepo
HOURGLASS_COMMIT=fc82a2b32c486bd6027288cf451ca20b8370784e

setup-hourglass: ## Clone and wire up hourglass-monorepo locally
	@echo "📦 Cloning hourglass-monorepo to $(HOURGLASS_DIR)..."
	@rm -rf $(HOURGLASS_DIR)
	@git clone $(GIT_HOURGLASS_REPO) $(HOURGLASS_DIR)
	@cd $(HOURGLASS_DIR) && git checkout $(HOURGLASS_COMMIT)
	@echo "🔧 Replacing Go module path to use local clone..."
	@go mod edit -replace=github.com/Layr-Labs/hourglass-monorepo=$(HOURGLASS_DIR)
	@go mod edit -replace=github.com/Layr-Labs/hourglass-monorepo/contracts=$(HOURGLASS_DIR)/contracts
	@go mod tidy
	@echo "✅ hourglass-monorepo (commit: $(HOURGLASS_COMMIT)) linked successfully"


build: deps setup-hourglass
	@mkdir -p $(OUT) || true
	@echo "Building binaries..."
	go build -o $(OUT)/performer ./avs/cmd/main.go

deps:
	GOPRIVATE=github.com/Layr-Labs/* go mod tidy


build/container:
	./.hourglass/scripts/buildContainer.sh


