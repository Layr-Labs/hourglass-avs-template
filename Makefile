
GO = $(shell which go)
OUT = ./bin

build: deps
	@mkdir -p $(OUT) || true
	@echo "Building binaries..."
	go build -o $(OUT)/performer ./avs/cmd/main.go

deps:
	go mod tidy


build/container:
	$(eval buildConfig := $(shell cat build.yaml))
	$(eval registry := $(shell yq '.container.registry' build.yaml))
	$(if $(registry),$(eval registry := $(registry)/),)

	docker build -t "$(registry)$(shell yq '.container.image' build.yaml):$(shell yq '.container.version' build.yaml)" .
