#!/usr/bin/env bash
set -e

# Build params
buildParams=$(cat ./.hourglass/build.yaml)
registry=$(echo "$buildParams" | yq -r '.container.registry')
image=$(echo "$buildParams" | yq -r '.container.image')

# Construct image name
if [ -n "$registry" ] && [ "$registry" != "null" ]; then
  fullImage="${registry}/${image}"
else
  fullImage="${image}"
fi

echo "Building container: $fullImage"

# Simple docker build 
docker build -t "$fullImage" .

# Get the image ID
IMAGE_ID=$(docker images --format "table {{.ID}}" --no-trunc "$fullImage" | tail -1)

echo "Built container: $fullImage"
echo "📋 Image ID: $IMAGE_ID"

# Export build info
echo "IMAGE_NAME=$fullImage" > /tmp/build_info
echo "IMAGE_ID=$IMAGE_ID" >> /tmp/build_info