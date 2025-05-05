set -o errexit -o nounset -o pipefail

# Navigate to script root
parent_path=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)
root_dir=$(realpath "$parent_path/../..")

cd "$root_dir/contracts"

# Convert hex to 256-bit decimal using Python
DECIMAL_KEY=$(python3 -c "print(int('$OPERATOR_PVT_KEY', 16))")
forge script script/RegisterOperatorToEigenLayer.s.sol:RegisterOperatorToEigenLayer --rpc-url "$RPC_URL" -vvvv  --broadcast --sig "run(uint)" -- "$DECIMAL_KEY" 
