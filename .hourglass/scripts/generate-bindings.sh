#!/bin/bash
set -e

# Paths
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$(dirname "$(dirname "$SCRIPT_DIR")")"

# source in helper functions
source "${PROJECT_ROOT}/.devkit/scripts/helpers/helpers.sh"

# Check if abigen is installed
ensureAbigen

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

error() {
    echo -e "${RED}[generate-bindings]${NC} $1" >&2
}
CONTRACTS_DIR="${PROJECT_ROOT}/contracts"
DEVKIT_CONTRACTS_DIR="${PROJECT_ROOT}/.devkit/contracts"
BINDING_DIR="${CONTRACTS_DIR}/bindings"
L1_CONTRACTS_DIR="${CONTRACTS_DIR}/src/l1-contracts"
L2_CONTRACTS_DIR="${CONTRACTS_DIR}/src/l2-contracts"

# Clean and recreate bindings directory
rm -rf "${BINDING_DIR}"
mkdir -p "${BINDING_DIR}"

# Function to generate binding for a contract
generate_binding() {
    local contract_name=$1
    local contract_type=$2  # l1 or l2
    local json_path="${DEVKIT_CONTRACTS_DIR}/out/${contract_name}.sol/${contract_name}.json"

    if [ ! -f "$json_path" ]; then
        error "Contract JSON not found: $json_path"
        error "Please run 'devkit avs build' first"
        return 1
    fi

    # Create a separate package directory for each contract
    local package_name=$(echo "${contract_name}" | tr '[:upper:]' '[:lower:]')
    local binding_out_dir="${BINDING_DIR}/${contract_type}/${package_name}"
    mkdir -p "${binding_out_dir}"

    # Extract ABI and bytecode
    cat "$json_path" | jq -r '.abi' > "${binding_out_dir}/tmp.abi"
    cat "$json_path" | jq -r '.bytecode.object' > "${binding_out_dir}/tmp.bin"

    # Generate Go binding
    log "Generating binding for ${contract_name} (${contract_type})..."
    abigen \
        --bin="${binding_out_dir}/tmp.bin" \
        --abi="${binding_out_dir}/tmp.abi" \
        --pkg="${package_name}" \
        --type="${contract_name}" \
        --out="${binding_out_dir}/${package_name}.go" \
        2>/dev/null

    if [ $? -eq 0 ]; then
        log "✓ Generated ${contract_type}/${package_name}/${package_name}.go"
    else
        error "Failed to generate binding for ${contract_name}"
    fi

    # Clean up temp files
    rm -f "${binding_out_dir}/tmp.abi" "${binding_out_dir}/tmp.bin"
}

log "Starting contract binding generation..."

# Generate bindings for L1 contracts
if [ -d "$L1_CONTRACTS_DIR" ]; then
    for contract_file in "$L1_CONTRACTS_DIR"/*.sol; do
        if [ -f "$contract_file" ]; then
            contract_name=$(basename "$contract_file" .sol)
            generate_binding "$contract_name" "l1"
        fi
    done
fi

# Generate bindings for L2 contracts
if [ -d "$L2_CONTRACTS_DIR" ]; then
    for contract_file in "$L2_CONTRACTS_DIR"/*.sol; do
        if [ -f "$contract_file" ]; then
            contract_name=$(basename "$contract_file" .sol)
            generate_binding "$contract_name" "l2"
        fi
    done
fi

log "Binding generation complete!"
log "Generated bindings are in: ${BINDING_DIR}/"