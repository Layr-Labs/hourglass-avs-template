set -o errexit -o nounset -o pipefail

# Navigate to script root
parent_path=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)
root_dir=$(realpath "$parent_path/../..")

cd "$root_dir/contracts"

DECIMAL_KEY=$(python3 -c "print(int('$OPERATOR_PVT_KEY', 16))")
forge script script/DepositIntoStrategies.s.sol:DepositIntoStrategies --rpc-url "$RPC_URL" -vvv  --broadcast --sig "run(address,uint)" -- "$STRATEGY_ADDRESS" "$DECIMAL_KEY"
