#!/usr/bin/env bash

buildParams=$(cat ./.hourglass/build.yaml)
registry=$(echo "$buildParams" | yq -r '.container.registry')
image=$(echo "$buildParams" | yq -r '.container.image')
tag=$(echo "$buildParams" | yq -r '.container.version')

if [[ ! -z "$registry" ]]; then
    image="$registry/$image"
fi

DOCKER_BUILDKIT=1 docker build \
    --ssh default \
    --progress=plain \
    -t "${image}:${tag}" .
