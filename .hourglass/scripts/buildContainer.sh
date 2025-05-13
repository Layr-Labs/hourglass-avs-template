#!/usr/bin/env bash

buildParams=$(cat ./.hourglass/build.yaml)
registry=$(echo "$buildParams" | yq -r '.container.registry')
image=$(echo "$buildParams" | yq -r '.container.image')
tag=$(echo "$buildParams" | yq -r '.container.version')

if [[ ! -z "$registry" ]]; then
    image="$registry/$image"
fi

# TODO(seanmcgary): remove this janky hack once the repo is public
mkdir -p ./.hourglass/.docker-build-tmp
cp -R ~/.ssh ./.hourglass/.docker-build-tmp

# check if ~/.gitconfig exists
if [[ -f ~/.gitconfig ]]; then
    cp ~/.gitconfig ./.hourglass/.docker-build-tmp/
fi

docker build -t "${image}:${tag}" .
