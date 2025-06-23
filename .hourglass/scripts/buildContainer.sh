#!/usr/bin/env bash
set -e

# Build params
buildParams=$(cat ./.hourglass/build.yaml)
registry=$(echo "$buildParams" | yq -r '.container.registry')
image=$(echo "$buildParams" | yq -r '.container.image')
tag=$(echo "$buildParams" | yq -r '.container.version')

# Construct full image name
if [ -z "$registry" ] || [ "$registry" = "null" ]; then
  fullImage="${image}:${tag}"
else
  fullImage="${registry}/${image}:${tag}"
fi

echo "Building container for current platform..."

docker build -t "$fullImage" .

echo "✅ Built container: $fullImage" 