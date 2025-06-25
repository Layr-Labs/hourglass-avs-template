#!/usr/bin/env bash
set -e

# Parse command line arguments for version support
CONFIG_DIR=""
VERSION_OVERRIDE=""

while [[ $# -gt 0 ]]; do
  case $1 in
    --config-dir)
      CONFIG_DIR="$2"
      shift 2
      ;;
    --version)
      VERSION_OVERRIDE="$2"
      shift 2
      ;;
    *)
      echo "Unknown option $1"
      exit 1
      ;;
  esac
done

# Read operator set mappings from devnet.yaml
echo "Reading operator set mappings from devnet.yaml..."

# Extract aggregator info
aggregator_operator_set_id=$(yq -r '.aggregator.operatorSetId' .hourglass/context/devnet.yaml)
aggregator_digest=$(yq -r '.aggregator.digest' .hourglass/context/devnet.yaml)
aggregator_registry_url=$(yq -r '.aggregator.registry_url' .hourglass/context/devnet.yaml)

# Extract executor info
executor_operator_set_id=$(yq -r '.executor.operatorSetId' .hourglass/context/devnet.yaml)
executor_digest=$(yq -r '.executor.digest' .hourglass/context/devnet.yaml)
executor_registry_url=$(yq -r '.executor.registry_url' .hourglass/context/devnet.yaml)

# Create operator set mapping as JSON
declare -A operator_set_mapping

# Create JSON struct for aggregator
aggregator_json=$(jq -n \
  --arg digest "$aggregator_digest" \
  --arg registry_url "$aggregator_registry_url" \
  '{digest: $digest, registry_url: $registry_url}')

# Create JSON struct for executor  
executor_json=$(jq -n \
  --arg digest "$executor_digest" \
  --arg registry_url "$executor_registry_url" \
  '{digest: $digest, registry_url: $registry_url}')

# Build the mapping
operator_set_mapping[$aggregator_operator_set_id]="[$aggregator_json]"
operator_set_mapping[$executor_operator_set_id]="[$executor_json]"

echo "Operator Set Mapping:"
echo "  OperatorSet $aggregator_operator_set_id (aggregator): ${operator_set_mapping[$aggregator_operator_set_id]}"
echo "  OperatorSet $executor_operator_set_id (executor): ${operator_set_mapping[$executor_operator_set_id]}"

# Function to get cached version
get_cached_version() {
  if [ -n "$CONFIG_DIR" ]; then
    local project_name=$(basename "$(pwd)")
    local version_cache_file="$CONFIG_DIR/versions/${project_name}_version"
    if [ -f "$version_cache_file" ]; then
      cat "$version_cache_file"
    else
      echo ""
    fi
  else
    echo ""
  fi
}

# Build params
buildParams=$(cat ./.hourglass/build.yaml)
registry=$(echo "$buildParams" | yq -r '.container.registry')
image=$(echo "$buildParams" | yq -r '.container.image')
config_tag=$(echo "$buildParams" | yq -r '.container.version')

# Determine final version/tag
if [ -n "$VERSION_OVERRIDE" ]; then
  tag="$VERSION_OVERRIDE"
else
  # Try to use cached version if available
  cached_version=$(get_cached_version)
  if [ -n "$cached_version" ]; then
    tag="$cached_version"
  else
    tag="$config_tag"
  fi
fi

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

# Create the final operator set mapping JSON output
operator_set_mapping_json=$(jq -n \
  --argjson operator_set_0 "${operator_set_mapping[$aggregator_operator_set_id]}" \
  --argjson operator_set_1 "${operator_set_mapping[$executor_operator_set_id]}" \
  '{
    ($aggregator_operator_set_id | tostring): $operator_set_0,
    ($executor_operator_set_id | tostring): $operator_set_1
  }')

# Export the result for the calling script
echo "IMAGE_CHANGED=$IMAGE_CHANGED" > /tmp/release_result
echo "ORIGINAL_IMAGE_ID=$ORIGINAL_IMAGE_ID" >> /tmp/release_result
echo "NEW_IMAGE_ID=$NEW_IMAGE_ID" >> /tmp/release_result
echo "OPERATOR_SET_MAPPING=$operator_set_mapping_json" >> /tmp/release_result

# Output the operator set mapping as JSON to stdout
echo "$operator_set_mapping_json"

# Return the boolean result (0 = no change, 1 = changed)
if [ "$IMAGE_CHANGED" = "true" ]; then
  exit 1
else
  exit 0
fi 