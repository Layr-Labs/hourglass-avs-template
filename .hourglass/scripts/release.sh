#!/usr/bin/env bash
set -e

# Multi-arch release script
# This script is invoked by `devkit avs release publish` for multi-architecture deployments
# It builds multi-platform images and creates an OCI compliant Image Index

# Usage: release.sh <registry_url> <image_name> <version> <local_image_id>

REGISTRY_URL="$1"
IMAGE_NAME="$2"
VERSION="$3"
LOCAL_IMAGE_ID="$4"

if [ -z "$REGISTRY_URL" ] || [ -z "$IMAGE_NAME" ] || [ -z "$VERSION" ] || [ -z "$LOCAL_IMAGE_ID" ]; then
    echo "Error: Missing required parameters"
    echo "Usage: release.sh <registry_url> <image_name> <version> <local_image_id>"
    exit 1
fi

# Construct image names
FULL_IMAGE_NAME="${REGISTRY_URL}/${IMAGE_NAME}:${VERSION}"

echo "Starting multi-arch release to registry..."
echo "  Registry: $REGISTRY_URL"
echo "  Image: ${IMAGE_NAME}:${VERSION}"
echo "  Local Image ID: $LOCAL_IMAGE_ID"

# Ensure builder is ready for multi-platform builds
if ! docker buildx inspect multiarch &>/dev/null; then
  echo "Creating multi-platform builder..."
  docker buildx create --name multiarch --driver docker-container --use
  docker buildx inspect --bootstrap
else
  docker buildx use multiarch
fi

# Build and push multi-platform image directly to registry
echo "Building and pushing multi-platform image..."
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --tag "$FULL_IMAGE_NAME" \
  --push \
  .

# Get the Image Index digest
echo "Getting Image Index digest..."
IMAGE_INDEX_DIGEST=$(docker buildx imagetools inspect "$FULL_IMAGE_NAME" --format '{{.Manifest.Digest}}')

echo "Multi-arch release completed successfully!"
echo "  Image Index: $FULL_IMAGE_NAME"
echo "  Digest: $IMAGE_INDEX_DIGEST"

# Verify OCI compliance
echo "Verifying OCI compliance..."
MANIFEST_CONTENT=$(docker buildx imagetools inspect "$FULL_IMAGE_NAME" --raw)
MEDIA_TYPE=$(echo "$MANIFEST_CONTENT" | jq -r '.mediaType')

if [[ "$MEDIA_TYPE" == "application/vnd.oci.image.index.v1+json" ]]; then
    echo "OCI compliant Image Index created"
    echo "   Media Type: $MEDIA_TYPE"
else
    echo "Warning: Media type is $MEDIA_TYPE (not OCI compliant)"
fi

# Show platform details
echo "Platform details:"
echo "$MANIFEST_CONTENT" | jq '.manifests[] | select(.platform) | {platform: .platform, digest: .digest}'

# Export the Image Index digest for devkit to use
echo "IMAGE_INDEX_DIGEST=$IMAGE_INDEX_DIGEST" > /tmp/release_digest
echo "FULL_IMAGE_NAME=$FULL_IMAGE_NAME" >> /tmp/release_digest

echo ""
echo "Release ready for ReleaseManager!"
echo "   Use digest: $IMAGE_INDEX_DIGEST" 