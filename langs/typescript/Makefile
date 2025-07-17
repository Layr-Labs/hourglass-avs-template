# -----------------------------------------------------------------------------
# This Makefile is used for building your AVS application.
#
# It contains basic targets for building the application, installing dependencies,
# and building a Docker container.
#
# Modify each target as needed to suit your application's requirements.
# -----------------------------------------------------------------------------

build: deps

deps:
	npm install

build/container:
	./.hourglass/scripts/buildContainer.sh

test: test-node test-forge

test-node:
	npm run test

test-forge:
	cd .devkit/contracts && forge test
