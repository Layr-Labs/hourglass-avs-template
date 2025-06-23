#!/usr/bin/env bash
set -e

# Ensure builder is ready
if ! docker buildx inspect multiarch &>/dev/null; then
  echo "[devkit] Creating multi-platform builder..."
  docker buildx create --name multiarch --driver docker-container --use
  docker buildx inspect --bootstrap
else
  docker buildx use multiarch
fi

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

echo "Building multi-platform images with individual IDs..."

# Build for linux/amd64
echo "Building linux/amd64..."
docker buildx build \
  --platform linux/amd64 \
  --tag "${fullImage}-amd64" \
  --load \
  .

# Build for linux/arm64  
echo "Building linux/arm64..."
docker buildx build \
  --platform linux/arm64 \
  --tag "${fullImage}-arm64" \
  --load \
  .

# Get image IDs for both platforms
AMD64_IMAGE_ID=$(docker images --format "{{.ID}}" --filter "reference=${fullImage}-amd64" | head -1)
ARM64_IMAGE_ID=$(docker images --format "{{.ID}}" --filter "reference=${fullImage}-arm64" | head -1)

echo "✅ Built multi-platform images:"
echo "  linux/amd64: ${fullImage}-amd64 (ID: $AMD64_IMAGE_ID)"
echo "  linux/arm64: ${fullImage}-arm64 (ID: $ARM64_IMAGE_ID)"

# Create OCI compliant multi-platform manifest
echo ""
echo "Creating OCI compliant multi-platform Image Index..."

# Build multi-platform with OCI output - keep files for easy inspection
OCI_TAR="./oci-manifest.tar"
OCI_DIR="./oci-manifest"

docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --tag "${fullImage}" \
  --output type=oci,dest="$OCI_TAR" \
  .

# Extract the tar for inspection
rm -rf "$OCI_DIR" 2>/dev/null || true
mkdir -p "$OCI_DIR"
tar -xf "$OCI_TAR" -C "$OCI_DIR"

# Extract the manifest digest from the OCI index
if [ -f "$OCI_DIR/index.json" ]; then
  # The index.json contains a reference to the actual manifest list
  # Get the digest of the actual manifest list (not the wrapper)
  ACTUAL_MANIFEST_DIGEST=$(cat "$OCI_DIR/index.json" | jq -r '.manifests[0].digest')
  
  # Find the actual manifest list file
  MANIFEST_LIST_FILE="$OCI_DIR/blobs/$(echo $ACTUAL_MANIFEST_DIGEST | cut -d':' -f1)/$(echo $ACTUAL_MANIFEST_DIGEST | cut -d':' -f2)"
  
  if [ -f "$MANIFEST_LIST_FILE" ]; then
    # This is the actual multi-platform manifest list with platform details
    MANIFEST_DIGEST="$ACTUAL_MANIFEST_DIGEST"
    
    echo "✅ Created OCI compliant multi-platform Image Index:"
    echo "  Manifest List: ${fullImage}"
    echo "  Digest: $MANIFEST_DIGEST"
    
    echo ""
    echo "=== OCI Multi-Platform Manifest List ==="
    cat "$MANIFEST_LIST_FILE" | jq .
  else
    echo "❌ Failed to find manifest list file"
    MANIFEST_DIGEST=""
  fi
else
  echo "❌ Failed to create OCI manifest"
  MANIFEST_DIGEST=""
fi

# Export all IDs for the build script to use
echo "AMD64_IMAGE_ID=$AMD64_IMAGE_ID" > /tmp/multiarch_image_ids
echo "ARM64_IMAGE_ID=$ARM64_IMAGE_ID" >> /tmp/multiarch_image_ids
echo "MANIFEST_DIGEST=$MANIFEST_DIGEST" >> /tmp/multiarch_image_ids
