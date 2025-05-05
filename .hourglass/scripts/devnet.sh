#!/bin/bash

set -o errexit -o nounset -o pipefail

# Navigate to script root
parent_path=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)
root_dir=$(realpath "$parent_path/../..")
cd "$root_dir"

RPC_URL=http://localhost:8545
PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# Initialize Git repo if not already
if [ ! -d .git ]; then
    echo "🔧 Initializing Git repository..."
    git init
fi

# Untrack submodules in Git index if already added
git rm --cached contracts/lib/eigenlayer-middleware 2>/dev/null || true
git rm --cached contracts/lib/forge-std 2>/dev/null || true
# Remove existing submodule folders
rm -rf contracts/lib/eigenlayer-middleware contracts/lib/forge-std

# Clean any submodule remnants
git config --remove-section submodule.contracts/lib/eigenlayer-middleware 2>/dev/null || true
git config --remove-section submodule.contracts/lib/forge-std 2>/dev/null || true
rm -rf .git/modules/contracts/lib/eigenlayer-middleware .git/modules/contracts/lib/forge-std

# Delete stale .gitmodules (must be staged or removed properly)
if [ -f .gitmodules ]; then
    echo "🧹 Removing stale .gitmodules"
    git rm --cached .gitmodules || true
    rm -f .gitmodules
fi

# # Re-add submodules
git submodule add https://github.com/Layr-Labs/eigenlayer-middleware contracts/lib/eigenlayer-middleware
git submodule add https://github.com/foundry-rs/forge-std contracts/lib/forge-std


# # Initialize submodules
git submodule update --init --recursive

cd "$root_dir/contracts"

cp .env.example .env
# # Call Forge script
# forge script script/DeployTaskMailbox.s.sol:DeployTaskMailbox --rpc-url "$RPC_URL" -vvvv --private-key "$PRIVATE_KEY" --broadcast --sig "run()"

# forge script script/DeployTaskAVSRegistrar.s.sol:DeployTaskAVSRegistrar --rpc-url "$RPC_URL" -vvvv --private-key "$PRIVATE_KEY" --broadcast --sig "run()" 

# forge script script/DeployAVSL2Contracts.s.sol:DeployAVSL2Contracts --rpc-url "$RPC_URL" -vvvv --broadcast --sig "run()" 

# forge script script/SetupAVSL1Contracts.s.sol:SetupAVSL1Contracts --rpc-url "$RPC_URL" -vvvv --broadcast --sig "run()" 

# forge script script/CreateOperatorSet.s.sol:CreateOperatorSet --rpc-url "$RPC_URL" -vvvv --broadcast --sig "run()" 
