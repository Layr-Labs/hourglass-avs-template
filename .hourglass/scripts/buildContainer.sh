#!/usr/bin/env bash
set -e

# Parse arguments
ARCHITECTURES=""
REGISTRY_URL=""

while [[ $# -gt 0 ]]; do
  case $1 in
    --architectures)
      ARCHITECTURES="$2"
      shift 2
      ;;
    --registry-url)
      REGISTRY_URL="$2"
      shift 2
      ;;
    *)
      echo "Unknown option $1"
      exit 1
      ;;
  esac
done

# Build params
buildParams=$(cat ./.hourglass/build.yaml)
registry=$(echo "$buildParams" | yq -r '.container.registry')
image=$(echo "$buildParams" | yq -r '.container.image')
tag=$(echo "$buildParams" | yq -r '.container.version')

# Use registry URL from parameter or config
if [ -n "$REGISTRY_URL" ]; then
  FINAL_REGISTRY="$REGISTRY_URL"
elif [ -n "$registry" ] && [ "$registry" != "null" ]; then
  FINAL_REGISTRY="$registry"
else
  FINAL_REGISTRY=""
fi

# Construct image names
if [ -n "$FINAL_REGISTRY" ]; then
  fullImage="${FINAL_REGISTRY}/${image}:${tag}"
else
  fullImage="${image}:${tag}"
fi

# Multi-arch build
if [ -n "$ARCHITECTURES" ]; then
  if [ -z "$FINAL_REGISTRY" ]; then
    echo "Error: Registry URL is required for multi-architecture builds"
    echo "Use --registry-url parameter or set registry in .hourglass/build.yaml"
    exit 1
  fi
  
  echo "Building multi-architecture image for: $ARCHITECTURES"
  echo "Pushing to registry: $FINAL_REGISTRY"
  
  # Setup buildx for multi-platform
  if ! docker buildx inspect multiarch &>/dev/null; then
    echo "Creating multi-platform builder..."
    docker buildx create --name multiarch --driver docker-container --use
    docker buildx inspect --bootstrap
  else
    docker buildx use multiarch
  fi
  
  # Build and push multi-arch
  docker buildx build \
    --platform "$ARCHITECTURES" \
    --tag "$fullImage" \
    --push \
    .
  
  echo "✅ Built and pushed multi-architecture image: $fullImage"
  echo "   Platforms: $ARCHITECTURES"
  
  # Get the Image Index digest and create a digest-based tag
  echo "Getting Image Index digest..."
  IMAGE_DIGEST=$(docker buildx imagetools inspect "$fullImage" | grep "Digest:" | head -1 | awk '{print $2}')
  DIGEST_TAG=$(echo "$IMAGE_DIGEST" | sed 's/sha256://')
  
  if [ -n "$FINAL_REGISTRY" ]; then
    fullImageWithDigestTag="${FINAL_REGISTRY}/${image}:${DIGEST_TAG}"
  else
    fullImageWithDigestTag="${image}:${DIGEST_TAG}"
  fi
  
  echo "📋 Image Index Digest: $IMAGE_DIGEST"
  echo "🏷️  Creating digest-based tag: $fullImageWithDigestTag"
  
  # Tag the image with digest-based tag
  docker buildx imagetools create "$fullImage" --tag "$fullImageWithDigestTag"
  
  echo "✅ Created additional digest-based tag: $fullImageWithDigestTag"
  
  # Export registry info for build script
  echo "MULTI_ARCH=true" > /tmp/build_info
  echo "IMAGE_NAME=$fullImage" >> /tmp/build_info
  echo "IMAGE_DIGEST_TAG=$fullImageWithDigestTag" >> /tmp/build_info
  echo "IMAGE_DIGEST=$IMAGE_DIGEST" >> /tmp/build_info
  echo "REGISTRY_URL=$FINAL_REGISTRY" >> /tmp/build_info
  
else
  # Single-arch build - create Image Index locally (no push)
  echo "Building container for current platform..."
  
  # Detect current platform
  CURRENT_PLATFORM="linux/$(uname -m | sed 's/x86_64/amd64/')"
  echo "Detected platform: $CURRENT_PLATFORM"
  
  # Setup buildx for consistency
  if ! docker buildx inspect multiarch &>/dev/null; then
    echo "Creating multi-platform builder..."
    docker buildx create --name multiarch --driver docker-container --use
    docker buildx inspect --bootstrap
  else
    docker buildx use multiarch
  fi
  
  # Build single-platform as Image Index (LOCAL ONLY - no --push)
  docker buildx build \
    --platform "$CURRENT_PLATFORM" \
    --tag "$fullImage" \
    --load \
    .
  
  echo "✅ Built single-platform Image Index locally: $fullImage"
  echo "   Platform: $CURRENT_PLATFORM"
  
  # Get the local Image Index digest
  echo "Getting local Image Index digest..."
  IMAGE_DIGEST=$(docker buildx imagetools inspect "$fullImage" | grep "Digest:" | head -1 | awk '{print $2}')
  
  echo "MULTI_ARCH=false" > /tmp/build_info
  echo "IMAGE_NAME=$fullImage" >> /tmp/build_info
  echo "IMAGE_DIGEST=$IMAGE_DIGEST" >> /tmp/build_info
  echo "REGISTRY_URL=$FINAL_REGISTRY" >> /tmp/build_info
fi