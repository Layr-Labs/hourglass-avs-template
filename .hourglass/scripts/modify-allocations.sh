#!/bin/bash
set -o errexit -o nounset -o pipefail

# Navigate to contract dir
parent_path=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)
root_dir=$(realpath "$parent_path/../..")
cd "$root_dir/contracts"

DECIMAL_KEY=$(python3 -c "print(int('$OPERATOR_PVT_KEY', 16))")
forge script script/ModifyAllocations.s.sol:ModifyAllocations --rpc-url "$RPC_URL" --broadcast -vvvv --sig "run(uint,address,uint64,address,uint32)" -- "$DECIMAL_KEY" "$STRATEGY_ADDRESS" "$ALLOCATION" "$AVS_ADDRESS" "$OPERATOR_SET_ID"
