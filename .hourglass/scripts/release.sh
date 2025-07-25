#!/usr/bin/env bash
set -e

# Parse command line arguments for version
VERSION=""
REGISTRY=""
IMAGE=""

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

if [ -z "$REGISTRY" ]; then
  echo "Error: --registry is required" >&2
  exit 1
fi

if [ -z "$IMAGE" ]; then
  echo "Error: --image is required" >&2
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

# Extract AVS name from config.yaml first
# The script is called from the project root, so config/config.yaml should be relative to that
avs_name=$(yq -r '.config.project.name' ./config/config.yaml 2>/dev/null || echo "")
if [ -z "$avs_name" ] || [ "$avs_name" = "null" ]; then
  echo "Error: Project name not found in config/config.yaml" >&2
  echo "Please ensure config.project.name is set in your config.yaml file" >&2
  echo "Current directory: $(pwd)" >&2
  echo "Looking for config at: ./config/config.yaml" >&2
  exit 1
fi

echo "AVS Name: ${avs_name}" >&2
echo "Registry: ${REGISTRY}" >&2

# Setup buildx for multi-platform builds
echo "Setting up multi-platform builder..." >&2
if ! docker buildx inspect multiarch >/dev/null 2>&1; then
  docker buildx create --name multiarch --driver docker-container --use >&2
  docker buildx inspect --bootstrap >&2
else
  docker buildx use multiarch >&2
fi

# Construct performer image name
# Use registry exactly as provided by the user
performer_full_image="${REGISTRY}:performer-${VERSION}"

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

# Generate runtime specs using gomplate with values approach
echo "Generating EigenRuntime specifications..." >&2

# Prepare context data for substitution
# Use registry exactly as provided
avs_repository="${REGISTRY}"

CONTEXT_DATA=$(jq -n \
  --arg avs_name "$avs_name" \
  --arg version "$VERSION" \
  --arg aggregator_registry "$aggregator_registry" \
  --arg aggregator_digest "$aggregator_digest" \
  --arg aggregator_operator_set_id "$aggregator_operator_set_id" \
  --arg executor_registry "$executor_registry" \
  --arg executor_digest "$executor_digest" \
  --arg executor_operator_set_id "$executor_operator_set_id" \
  --arg performer_registry "$avs_repository" \
  --arg performer_digest "$performer_digest" \
  '{
    AVS_NAME: $avs_name,
    VERSION: $version,
    AGGREGATOR_REGISTRY: $aggregator_registry,
    AGGREGATOR_DIGEST: $aggregator_digest,
    AGGREGATOR_OPERATOR_SET_ID: $aggregator_operator_set_id,
    EXECUTOR_REGISTRY: $executor_registry,
    EXECUTOR_DIGEST: $executor_digest,
    EXECUTOR_OPERATOR_SET_ID: $executor_operator_set_id,
    PERFORMER_REGISTRY: $performer_registry,
    PERFORMER_DIGEST: $performer_digest
  }')

# Determine which values file to use based on context
CONTEXT_NAME="${CONTEXT_NAME:-devnet}"
USER_VALUES_FILE="specs/runtime/${CONTEXT_NAME}.yaml"

# Merge values
MERGED_VALUES=$(mktemp).yaml
bash .hourglass/scripts/merge-values.sh \
  ".hourglass/spec/defaults.yaml" \
  "$USER_VALUES_FILE" \
  "$CONTEXT_DATA" \
  "$MERGED_VALUES"

# Generate aggregator runtime spec
echo "Generating aggregator runtime spec with context: $CONTEXT_NAME" >&2
aggregator_runtime_spec=$(gomplate -d values=file://"$MERGED_VALUES" -f .hourglass/spec/aggregator-runtime-template.yaml 2>/dev/null)
gomplate_exit_code=$?
if [ $gomplate_exit_code -ne 0 ]; then
  echo "Error: gomplate failed to generate aggregator runtime spec (exit code: $gomplate_exit_code)" >&2
  echo "Running gomplate with debug output:" >&2
  gomplate -d values=file://"$MERGED_VALUES" -f .hourglass/spec/aggregator-runtime-template.yaml
  rm -f "$MERGED_VALUES"
  exit 1
fi

# Generate executor runtime spec
echo "Generating executor runtime spec with context: $CONTEXT_NAME" >&2
executor_runtime_spec=$(gomplate -d values=file://"$MERGED_VALUES" -f .hourglass/spec/executor-runtime-template.yaml 2>/dev/null)
gomplate_exit_code=$?
if [ $gomplate_exit_code -ne 0 ]; then
  echo "Error: gomplate failed to generate executor runtime spec (exit code: $gomplate_exit_code)" >&2
  echo "Running gomplate with debug output:" >&2
  gomplate -d values=file://"$MERGED_VALUES" -f .hourglass/spec/executor-runtime-template.yaml
  rm -f "$MERGED_VALUES"
  exit 1
fi

# Clean up
rm -f "$MERGED_VALUES"

# Create JSON structs for operator sets with runtime specs
aggregator_json=$(jq -n \
  --arg digest "$aggregator_digest" \
  --arg registry "$aggregator_registry" \
  --arg runtimeSpec "$aggregator_runtime_spec" \
  '{digest: $digest, registry: $registry, runtimeSpec: $runtimeSpec}')

executor_json=$(jq -n \
  --arg digest "$executor_digest" \
  --arg registry "$executor_registry" \
  --arg runtimeSpec "$executor_runtime_spec" \
  '{digest: $digest, registry: $registry, runtimeSpec: $runtimeSpec}')

# Create the final operator set mapping JSON output with runtime specs
operator_set_mapping_json=$(jq -n \
  --arg agg_id "$aggregator_operator_set_id" \
  --argjson agg_data "$aggregator_json" \
  --arg exec_id "$executor_operator_set_id" \
  --argjson exec_data "$executor_json" \
  '{
    ($agg_id): $agg_data,
    ($exec_id): $exec_data
  }')

# Output the operator set mapping to stdout
echo "$operator_set_mapping_json"