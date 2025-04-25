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

deps:
	go mod tidy


build/container:
	$(eval buildConfig := $(shell cat ./.hourglass/build.yaml))
	$(eval registry := $(shell cat ./.hourglass/build.yaml | yq -r '.container.registry'))
	$(if $(registry),$(eval registry := $(registry)/),)

	docker build -t "$(registry)$(shell cat ./.hourglass/build.yaml | yq -r '.container.image'):$(shell cat ./.hourglass/build.yaml | yq -r '.container.version')" .
