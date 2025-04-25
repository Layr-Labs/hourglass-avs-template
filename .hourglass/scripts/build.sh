#!/usr/bin/env bash

BUILD_CONTAINER=${BUILD_CONTAINER:-"false"}

if [[ "$BUILD_CONTAINER" == "true" ]]; then
    make build/container
fi

make build
