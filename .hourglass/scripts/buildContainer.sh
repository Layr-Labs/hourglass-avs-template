#!/usr/bin/env bash
set -e

buildParams=$(cat ./.hourglass/build.yaml)
registry=$(echo "$buildParams" | yq -r '.container.registry')
image=$(echo "$buildParams" | yq -r '.container.image')
tag=$(echo "$buildParams" | yq -r '.container.version')

if [[ ! -z "$registry" ]]; then
    image="$registry/$image"
fi

# Build single multi-platform OCI image index (creates one image ID for all platforms)
echo "Building multi-platform OCI image index for: linux/amd64, linux/arm64"

docker buildx build \
    --platform linux/amd64,linux/arm64 \
    --progress=plain \
    --tag "${image}:${tag}" \
    .

echo "Multi-platform OCI image index built successfully: ${image}:${tag}"
echo "This single image ID contains all platform variants (linux/amd64, linux/arm64)"
