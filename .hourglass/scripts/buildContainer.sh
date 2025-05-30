#!/usr/bin/env bash
set -e

buildParams=$(cat ./.hourglass/build.yaml)
registry=$(echo "$buildParams" | yq -r '.container.registry')
image=$(echo "$buildParams" | yq -r '.container.image')
tag=$(echo "$buildParams" | yq -r '.container.version')

if [[ ! -z "$registry" ]]; then
    image="$registry/$image"
fi

docker build \
    --progress=plain \
    -t "${image}:${tag}" .
