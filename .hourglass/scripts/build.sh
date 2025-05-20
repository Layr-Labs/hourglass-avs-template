#!/usr/bin/env bash
set -e

BUILD_CONTAINER=${BUILD_CONTAINER:-"false"}

if [[ "$BUILD_CONTAINER" == "true" ]]; then
    ./.hourglass/scripts/buildContainer.sh
fi

make build
