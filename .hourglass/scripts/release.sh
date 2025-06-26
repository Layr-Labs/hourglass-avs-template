#!/usr/bin/env bash
set -e

# Parse command line arguments for version
VERSION=""
REGISTRY_URL=""

while [[ $# -gt 0 ]]; do
  case $1 in
    --version)
      VERSION="$2"
      shift 2
      ;;
    --registry-url)
      REGISTRY_URL="$2"
      shift 2
      ;;
    *)
      echo "Unknown option $1" >&2
      exit 1
      ;;
  esac
done

# Ensure version is provided
if [ -z "$VERSION" ]; then
  echo "Error: --version is required" >&2
  exit 1
fi

# Read operator set mappings from devnet.yaml
echo "Reading operator set mappings from devnet.yaml..." >&2

# Extract aggregator info
aggregator_operator_set_id=$(yq -r '.aggregator.operatorSetId' .hourglass/context/devnet.yaml)
aggregator_digest=$(yq -r '.aggregator.digest' .hourglass/context/devnet.yaml)
aggregator_registry_url=$(yq -r '.aggregator.registry_url' .hourglass/context/devnet.yaml)

# Extract executor info
executor_operator_set_id=$(yq -r '.executor.operatorSetId' .hourglass/context/devnet.yaml)
executor_digest=$(yq -r '.executor.digest' .hourglass/context/devnet.yaml)
executor_registry_url=$(yq -r '.executor.registry_url' .hourglass/context/devnet.yaml)

# Create JSON structs for operator sets
aggregator_json=$(jq -n \
  --arg digest "$aggregator_digest" \
  --arg registry_url "$aggregator_registry_url" \
  '{digest: $digest, registry_url: $registry_url}')

executor_json=$(jq -n \
  --arg digest "$executor_digest" \
  --arg registry_url "$executor_registry_url" \
  '{digest: $digest, registry_url: $registry_url}')

echo "Operator Set Mapping:" >&2
echo "  OperatorSet $aggregator_operator_set_id (aggregator): [$aggregator_json]" >&2
echo "  OperatorSet $executor_operator_set_id (executor): [$executor_json]" >&2


# Build params
buildParams=$(cat ./.hourglass/build.yaml)
image=$(echo "$buildParams" | yq -r '.container.image')

# Use provided version
tag="$VERSION"

# Construct original and temporary image names
if [ -n "$registry" ] && [ "$registry" != "null" ]; then
  originalImage="${registry}/${image}"
  tempImage="${registry}/${image}-release"
else
  originalImage="${image}"
  tempImage="${image}-release"
fi

echo "Rebuilding container for release comparison..." >&2
echo "Original image: $originalImage" >&2
echo "Temporary image: $tempImage" >&2

# Get the original image ID (from the build)
ORIGINAL_IMAGE_ID=$(docker images --format "table {{.ID}}" --no-trunc "$originalImage" | tail -1)

if [ -z "$ORIGINAL_IMAGE_ID" ]; then
  echo "Error: Original image $originalImage not found. Run 'devkit avs build' first." >&2
  exit 1
fi

echo "Original Image ID: $ORIGINAL_IMAGE_ID" >&2

# Build with temporary tag
docker build -t "$tempImage" . >&2

# Get the new image ID
NEW_IMAGE_ID=$(docker images --format "table {{.ID}}" --no-trunc "$tempImage" | tail -1)

echo "New Image ID: $NEW_IMAGE_ID" >&2

# Compare image IDs and set result
if [ "$ORIGINAL_IMAGE_ID" = "$NEW_IMAGE_ID" ]; then
  echo "Image unchanged - no rebuild needed" >&2
  echo "Both images have the same ID: $ORIGINAL_IMAGE_ID" >&2
  IMAGE_CHANGED="false"
else
  echo "Image changed - rebuild detected" >&2
  echo "Original: $ORIGINAL_IMAGE_ID" >&2
  echo "New:      $NEW_IMAGE_ID" >&2
  IMAGE_CHANGED="true"
fi

# Clean up temporary image
docker rmi "$tempImage" >/dev/null 2>&1 || true

# Setup buildx for multi-platform builds
echo "Setting up multi-platform builder..." >&2
if ! docker buildx inspect multiarch >/dev/null 2>&1; then
  docker buildx create --name multiarch --driver docker-container --use >&2
  docker buildx inspect --bootstrap >&2
else
  docker buildx use multiarch >&2
fi

# Build and push multi-arch performer image
project_name=$(basename "$(pwd)")
performer_image_name="${project_name}-performer-op-set-1"

# Construct performer image name based on registry presence
if [ -n "$REGISTRY_URL" ]; then
  performer_full_image="${REGISTRY_URL}/${performer_image_name}:${tag}"
else
  performer_full_image="${performer_image_name}:${tag}"
fi

echo "Building multi-architecture performer image: ${performer_full_image}" >&2
echo "Platforms: linux/amd64,linux/arm64" >&2

# Build and push multi-arch
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --tag "$performer_full_image" \
  --push \
  . >&2

# Get the Image Index digest
echo "Getting Image Index digest..." >&2
performer_digest=$(docker buildx imagetools inspect "$performer_full_image" | grep "Digest:" | head -n 1 | awk '{print $2}')
if [ -z "$performer_digest" ]; then
  echo "Error: Could not get performer image digest" >&2
  exit 1
fi
echo "📋 Performer Image Index Digest: $performer_digest" >&2

# Create performer JSON
performer_json=$(jq -n \
  --arg digest "$performer_digest" \
  --arg registry_url "$REGISTRY_URL" \
  '{digest: $digest, registry_url: $registry_url}')

# Create the final operator set mapping JSON output with performer
operator_set_mapping_json=$(jq -n \
  --arg agg_id "$aggregator_operator_set_id" \
  --argjson agg_data "[$aggregator_json]" \
  --arg exec_id "$executor_operator_set_id" \
  --argjson exec_data "[$executor_json, $performer_json]" \
  '{
    ($agg_id): $agg_data,
    ($exec_id): $exec_data
  }')

# Output the operator set mapping to stdout
echo "$operator_set_mapping_json"

# Export build results to file
echo "IMAGE_CHANGED=$IMAGE_CHANGED" > /tmp/release_result
echo "ORIGINAL_IMAGE_ID=$ORIGINAL_IMAGE_ID" >> /tmp/release_result
echo "NEW_IMAGE_ID=$NEW_IMAGE_ID" >> /tmp/release_result 

# Return the boolean result (0 = no change, 1 = changed)
if [ "$IMAGE_CHANGED" = "true" ]; then
  exit 1
else
  exit 0
fi 