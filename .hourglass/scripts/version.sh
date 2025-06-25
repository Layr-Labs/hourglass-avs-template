#!/usr/bin/env bash
set -e

# Version utility for hourglass-avs-template
# Supports semantic versioning with automatic increment
# Stores version in devkit config directory to survive template upgrades

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BUILD_YAML_PATH="$SCRIPT_DIR/../build.yaml"

# Default values
ACTION="get"
BUMP_TYPE="patch"
NEW_VERSION=""
CONFIG_DIR=""
USE_CACHE=true

# Usage function
usage() {
  cat << EOF
Version Management Utility for Hourglass AVS

USAGE:
  ./version.sh [OPTIONS] [--config-dir <dir>]

ACTIONS:
  get                 Get current version (default)
  set <version>       Set specific version (e.g., 1.2.3)
  bump <type>         Bump version by type: major, minor, patch (default: patch)
  git-tag             Use latest git tag as version
  commit              Use short git commit hash as version
  timestamp           Use timestamp as version (YYYY.MM.DD-HHMM)
  sync                Sync cached version to build.yaml
  reset               Reset to default version (0.1.0)

OPTIONS:
  --help, -h          Show this help message
  --dry-run           Show what would be changed without making changes
  --use-cache         Always use cached version (survives upgrades)
  --no-cache          Don't use cache, work directly with build.yaml
  --config-dir <dir>  Config directory

EXAMPLES:
  ./version.sh --config-dir ~/.config/devkit                    # Get current version
  ./version.sh --config-dir ~/.config/devkit set 1.0.0         # Set version to 1.0.0
  ./version.sh --config-dir ~/.config/devkit bump major        # Bump major version
  ./version.sh --config-dir ~/.config/devkit bump minor        # Bump minor version
  ./version.sh --config-dir ~/.config/devkit bump patch        # Bump patch version
  ./version.sh --config-dir ~/.config/devkit git-tag           # Use latest git tag
  ./version.sh --config-dir ~/.config/devkit commit            # Use git commit hash
  ./version.sh --config-dir ~/.config/devkit timestamp         # Use current timestamp
  ./version.sh --config-dir ~/.config/devkit sync              # Sync cached version to build.yaml
  ./version.sh --config-dir ~/.config/devkit reset             # Reset to default version

CACHE:
  Version is cached in <config-dir>/versions/<project-name>_version
  This ensures your version persists through template upgrades.

EOF
}

# Parse command line arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    get)
      ACTION="get"
      shift
      ;;
    set)
      ACTION="set"
      NEW_VERSION="$2"
      shift 2
      ;;
    bump)
      ACTION="bump"
      BUMP_TYPE="${2:-patch}"
      shift 2
      ;;
    git-tag)
      ACTION="git-tag"
      shift
      ;;
    commit)
      ACTION="commit"
      shift
      ;;
    timestamp)
      ACTION="timestamp"
      shift
      ;;
    sync)
      ACTION="sync"
      shift
      ;;
    reset)
      ACTION="reset"
      shift
      ;;
    --dry-run)
      DRY_RUN=true
      shift
      ;;
    --use-cache)
      USE_CACHE=true
      shift
      ;;
    --no-cache)
      USE_CACHE=false
      shift
      ;;
    --config-dir)
      CONFIG_DIR="$2"
      shift 2
      ;;
    --help|-h)
      usage
      exit 0
      ;;
    *)
      echo "Unknown option: $1"
      usage
      exit 1
      ;;
  esac
done

# Set up version cache file path
PROJECT_NAME=$(basename "$(pwd)")
VERSION_CACHE_FILE="$CONFIG_DIR/versions/${PROJECT_NAME}_version"
# Ensure cache directory exists
mkdir -p "$(dirname "$VERSION_CACHE_FILE")"

# Function to get cached version
get_cached_version() {
  if [ -f "$VERSION_CACHE_FILE" ]; then
    cat "$VERSION_CACHE_FILE"
  else
    echo ""
  fi
}

# Function to set cached version
set_cached_version() {
  local version="$1"
  echo "$version" > "$VERSION_CACHE_FILE"
  echo "Version cached: $version (in $VERSION_CACHE_FILE)"
}

# Function to get version from build.yaml
get_build_yaml_version() {
  if [ ! -f "$BUILD_YAML_PATH" ]; then
    echo "Error: build.yaml not found at $BUILD_YAML_PATH"
    exit 1
  fi
  
  yq -r '.container.version' "$BUILD_YAML_PATH"
}

# Function to get current version (with cache logic)
get_current_version() {
  local cached_version=$(get_cached_version)
  local build_yaml_version=$(get_build_yaml_version)
  
  if [ "$USE_CACHE" = true ] && [ -n "$cached_version" ]; then
    echo "$cached_version"
  else
    echo "$build_yaml_version"
  fi
}

# Function to validate semantic version
validate_semver() {
  local version="$1"
  if [[ ! "$version" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Error: Invalid semantic version format. Expected: MAJOR.MINOR.PATCH (e.g., 1.2.3)"
    exit 1
  fi
}

# Function to bump semantic version
bump_version() {
  local current_version="$1"
  local bump_type="$2"
  
  validate_semver "$current_version"
  
  IFS='.' read -r major minor patch <<< "$current_version"
  
  case "$bump_type" in
    major)
      echo "$((major + 1)).0.0"
      ;;
    minor)
      echo "${major}.$((minor + 1)).0"
      ;;
    patch)
      echo "${major}.${minor}.$((patch + 1))"
      ;;
    *)
      echo "Error: Invalid bump type. Use: major, minor, or patch"
      exit 1
      ;;
  esac
}

# Function to get git tag
get_git_tag() {
  local git_version=$(git describe --tags --abbrev=0 2>/dev/null || echo "")
  if [ -n "$git_version" ]; then
    echo "$git_version"
  else
    echo "Error: No git tags found"
    exit 1
  fi
}

# Function to get git commit hash
get_git_commit() {
  git rev-parse --short HEAD
}

# Function to get timestamp
get_timestamp() {
  date +%Y.%m.%d-%H%M
}

# Function to update version in build.yaml
update_build_yaml() {
  local new_version="$1"
  
  if [ "$DRY_RUN" = true ]; then
    echo "[DRY RUN] Would update build.yaml version to: $new_version"
    return
  fi
  
  # Create backup
  cp "$BUILD_YAML_PATH" "$BUILD_YAML_PATH.backup"
  
  # Update version
  yq -i ".container.version = \"$new_version\"" "$BUILD_YAML_PATH"
  
  echo "build.yaml updated to: $new_version"
  echo "Backup saved: $BUILD_YAML_PATH.backup"
}

# Function to update version (both cache and build.yaml)
update_version() {
  local new_version="$1"
  
  if [ "$DRY_RUN" = true ]; then
    echo "[DRY RUN] Would update version to: $new_version"
    echo "[DRY RUN] Would cache version in: $VERSION_CACHE_FILE"
    echo "[DRY RUN] Would update build.yaml: $BUILD_YAML_PATH"
    return
  fi
  
  # Always update cache
  set_cached_version "$new_version"
  
  # Update build.yaml
  update_build_yaml "$new_version"
}

# Function to sync cached version to build.yaml
sync_version() {
  local cached_version=$(get_cached_version)
  
  if [ -z "$cached_version" ]; then
    echo "No cached version found. Nothing to sync."
    return
  fi
  
  echo "Syncing cached version ($cached_version) to build.yaml"
  update_build_yaml "$cached_version"
}

# Main logic
case "$ACTION" in
  get)
    current_version=$(get_current_version)
    cached_version=$(get_cached_version)
    build_yaml_version=$(get_build_yaml_version)
    
    echo "Current version: $current_version"
    
    if [ "$USE_CACHE" = true ]; then
      echo "Cached version: ${cached_version:-"none"}"
      echo "build.yaml version: $build_yaml_version"
      echo "Cache file: $VERSION_CACHE_FILE"
    fi
    ;;
  set)
    if [ -z "$NEW_VERSION" ]; then
      echo "Error: No version specified"
      usage
      exit 1
    fi
    validate_semver "$NEW_VERSION"
    update_version "$NEW_VERSION"
    ;;
  bump)
    current_version=$(get_current_version)
    new_version=$(bump_version "$current_version" "$BUMP_TYPE")
    echo "Bumping $BUMP_TYPE version: $current_version -> $new_version"
    update_version "$new_version"
    ;;
  git-tag)
    git_version=$(get_git_tag)
    echo "Using git tag version: $git_version"
    update_version "$git_version"
    ;;
  commit)
    commit_version=$(get_git_commit)
    echo "Using git commit version: $commit_version"
    update_version "$commit_version"
    ;;
  timestamp)
    timestamp_version=$(get_timestamp)
    echo "Using timestamp version: $timestamp_version"
    update_version "$timestamp_version"
    ;;
  sync)
    sync_version
    ;;
  reset)
    echo "Resetting to default version: 0.1.0"
    update_version "0.1.0"
    ;;
esac 