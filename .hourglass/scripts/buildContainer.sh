#!/usr/bin/env bash
set -e

# Parse command line arguments
TAG=""
REGISTRY=""
IMAGE=""

while [[ $# -gt 0 ]]; do
  case $1 in
    --tag)
      TAG="$2"
      shift 2
      ;;
    --registry)
      REGISTRY="$2"
      shift 2
      ;;
    --image)
      IMAGE="$2"
      shift 2
      ;;
    *)
      echo "Unknown option $1" >&2
      exit 1
      ;;
  esac
done

# Ensure required arguments are provided
if [ -z "$IMAGE" ]; then
  echo "Error: --image is required" >&2
  exit 1
fi

# Construct image name
if [ -n "$REGISTRY" ] && [ "$REGISTRY" != "null" ]; then
  fullImage="${REGISTRY}/${IMAGE}"
else
  fullImage="${IMAGE}"
fi

# Add tag if provided
if [ -n "$TAG" ]; then
  fullImage="${fullImage}:${TAG}"
fi

echo "Building container: $fullImage" >&2

# Simple docker build 
docker build -t "$fullImage" . >&2

# Get the image ID
IMAGE_ID=$(docker images --format "table {{.ID}}" --no-trunc "$fullImage" | tail -1)

echo "Built container: $fullImage" >&2
echo "📋 Image ID: $IMAGE_ID" >&2

# Output build info as JSON to stdout
jq -n \
  --arg image "$fullImage" \
  --arg image_id "$IMAGE_ID" \
  '{image: $image, image_id: $image_id}'