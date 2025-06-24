#!/usr/bin/env bash
set -e

# Build params
buildParams=$(cat ./.hourglass/build.yaml)
registry=$(echo "$buildParams" | yq -r '.container.registry')
image=$(echo "$buildParams" | yq -r '.container.image')
tag=$(echo "$buildParams" | yq -r '.container.version')

# Construct original and temporary image names
if [ -n "$registry" ] && [ "$registry" != "null" ]; then
  originalImage="${registry}/${image}:${tag}"
  tempImage="${registry}/${image}:${tag}-release"
else
  originalImage="${image}:${tag}"
  tempImage="${image}:${tag}-release"
fi

echo "Rebuilding container for release comparison..."
echo "Original image: $originalImage"
echo "Temporary image: $tempImage"

# Get the original image ID (from the build)
ORIGINAL_IMAGE_ID=$(docker images --format "table {{.ID}}" --no-trunc "$originalImage" | tail -1)

if [ -z "$ORIGINAL_IMAGE_ID" ]; then
  echo "Error: Original image $originalImage not found. Run 'devkit avs build' first."
  exit 1
fi

echo "Original Image ID: $ORIGINAL_IMAGE_ID"

# Build with temporary tag
docker build -t "$tempImage" .

# Get the new image ID
NEW_IMAGE_ID=$(docker images --format "table {{.ID}}" --no-trunc "$tempImage" | tail -1)

echo "New Image ID: $NEW_IMAGE_ID"

# Compare image IDs and set result
if [ "$ORIGINAL_IMAGE_ID" = "$NEW_IMAGE_ID" ]; then
  echo "Image unchanged - no rebuild needed"
  echo "Both images have the same ID: $ORIGINAL_IMAGE_ID"
  IMAGE_CHANGED="false"
else
  echo "Image changed - rebuild detected"
  echo "Original: $ORIGINAL_IMAGE_ID"
  echo "New:      $NEW_IMAGE_ID"
  IMAGE_CHANGED="true"
fi

# Clean up temporary image
docker rmi "$tempImage" || true

# Export the result for the calling script
echo "IMAGE_CHANGED=$IMAGE_CHANGED" > /tmp/release_result
echo "ORIGINAL_IMAGE_ID=$ORIGINAL_IMAGE_ID" >> /tmp/release_result
echo "NEW_IMAGE_ID=$NEW_IMAGE_ID" >> /tmp/release_result

# Return the boolean result (0 = no change, 1 = changed)
if [ "$IMAGE_CHANGED" = "true" ]; then
  exit 1
else
  exit 0
fi 