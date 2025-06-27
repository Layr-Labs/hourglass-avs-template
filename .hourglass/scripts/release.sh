#!/usr/bin/env bash
set -e

# Parse command line arguments for version
VERSION=""
REGISTRY=""
IMAGE=""
ORIGINAL_IMAGE_ID=""

while [[ $# -gt 0 ]]; do
  case $1 in
    --version)
      VERSION="$2"
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
    --original-image-id)
      ORIGINAL_IMAGE_ID="$2"
      shift 2
      ;;
    *)
      echo "Unknown option $1" >&2
      exit 1
      ;;
  esac
done

# Ensure required arguments are provided
if [ -z "$VERSION" ]; then
  echo "Error: --version is required" >&2
  exit 1
fi

if [ -z "$IMAGE" ]; then
  echo "Error: --image is required" >&2
  exit 1
fi

if [ -z "$ORIGINAL_IMAGE_ID" ]; then
  echo "Error: --original-image-id is required" >&2
  exit 1
fi


# Read operator set mappings from devnet.yaml
echo "Reading operator set mappings from devnet.yaml..." >&2

# Extract aggregator info
aggregator_operator_set_id=$(yq -r '.aggregator.operatorSetId' .hourglass/context/devnet.yaml)
aggregator_digest=$(yq -r '.aggregator.digest' .hourglass/context/devnet.yaml)
aggregator_registry=$(yq -r '.aggregator.registry' .hourglass/context/devnet.yaml)

# Extract executor info
executor_operator_set_id=$(yq -r '.executor.operatorSetId' .hourglass/context/devnet.yaml)
executor_digest=$(yq -r '.executor.digest' .hourglass/context/devnet.yaml)
executor_registry=$(yq -r '.executor.registry' .hourglass/context/devnet.yaml)

# Create JSON structs for operator sets
aggregator_json=$(jq -n \
  --arg digest "$aggregator_digest" \
  --arg registry "$aggregator_registry" \
  '{digest: $digest, registry: $registry}')

executor_json=$(jq -n \
  --arg digest "$executor_digest" \
  --arg registry "$executor_registry" \
  '{digest: $digest, registry: $registry}')

echo "Operator Set Mapping:" >&2
echo "  OperatorSet $aggregator_operator_set_id (aggregator): [$aggregator_json]" >&2
echo "  OperatorSet $executor_operator_set_id (executor): [$executor_json]" >&2

# Build temporary image for comparison
echo "Building temporary image for comparison..." >&2
build_cmd=".hourglass/scripts/buildContainer.sh --image $IMAGE" # We are not giving registry url intentionally. Since its not pushing to registry , we just want to compare
temp_build=$(bash -c "$build_cmd --tag ${VERSION}-release")
NEW_IMAGE_ID=$(echo "$temp_build" | jq -r '.image_id')
tempImage=$(echo "$temp_build" | jq -r '.image')

echo "Original Image ID: $ORIGINAL_IMAGE_ID" >&2
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

# if image changed , return 1 and inform the user they need to  build again before performing a release.
if [ "$IMAGE_CHANGED" = "true" ]; then
  echo "Image changed - rebuild detected" >&2
  echo "Original: $ORIGINAL_IMAGE_ID" >&2
  echo "New:      $NEW_IMAGE_ID" >&2
  echo "Please run build again before performing a release." >&2
  exit 1
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

# Construct performer image name based on registry presence
if [ -n "$registry" ]; then
  performer_full_image="${registry}/${IMAGE}:${VERSION}"
else
  performer_full_image="${IMAGE}:${VERSION}"
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
echo "Performer Image Index Digest: $performer_digest" >&2

# Create performer JSON
performer_json=$(jq -n \
  --arg digest "$performer_digest" \
  --arg registry "$registry" \
  '{digest: $digest, registry: $registry}')

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
