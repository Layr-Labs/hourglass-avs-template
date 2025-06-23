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

# Create the release manifest directly as an OCI layout
echo ""
echo "Creating Release Manifest OCI layout..."

RELEASE_OCI_DIR="./release-manifest"
rm -rf "$RELEASE_OCI_DIR" 2>/dev/null || true

# Build multi-platform image to tar, then extract to directory
RELEASE_OCI_TAR="./release-manifest.tar"
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --tag "${fullImage}" \
  --output type=oci,dest="$RELEASE_OCI_TAR" \
  .

# Extract the tar to create the OCI layout directory
mkdir -p "$RELEASE_OCI_DIR"
tar -xf "$RELEASE_OCI_TAR" -C "$RELEASE_OCI_DIR"

# Get the container manifest digest from the OCI layout
CONTAINER_MANIFEST_DIGEST=""
if [ -f "$RELEASE_OCI_DIR/index.json" ]; then
  CONTAINER_MANIFEST_DIGEST=$(cat "$RELEASE_OCI_DIR/index.json" | jq -r '.manifests[0].digest')
  echo "✅ Created Release Manifest OCI layout:"
  echo "  Container Image: ${fullImage}"
  echo "  Container Digest: $CONTAINER_MANIFEST_DIGEST"
  
  # Clean up the tar file
  rm -f "$RELEASE_OCI_TAR"
else
  echo "❌ Failed to create Release Manifest OCI layout"
  exit 1
fi

# Export all IDs and digests for the build script to use
echo "AMD64_IMAGE_ID=$AMD64_IMAGE_ID" > /tmp/multiarch_image_ids
echo "ARM64_IMAGE_ID=$ARM64_IMAGE_ID" >> /tmp/multiarch_image_ids
echo "CONTAINER_MANIFEST_DIGEST=$CONTAINER_MANIFEST_DIGEST" >> /tmp/multiarch_image_ids

# For publishing, use the container manifest digest
MANIFEST_DIGEST="$CONTAINER_MANIFEST_DIGEST"
echo "MANIFEST_DIGEST=$MANIFEST_DIGEST" >> /tmp/multiarch_image_ids

echo ""
echo "🎉 Build completed successfully!"
echo "   Container Image Digest: $CONTAINER_MANIFEST_DIGEST"
echo "   Use Container Manifest Digest for publishing: $CONTAINER_MANIFEST_DIGEST" 