#!/usr/bin/env bash
set -e

# Parse command line arguments for version support only
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

# Construct image name
if [ -n "$registry" ] && [ "$registry" != "null" ]; then
  fullImage="${registry}/${image}:${tag}"
else
  fullImage="${image}:${tag}"
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
echo "VERSION=$tag" >> /tmp/build_info