set -o errexit -o nounset -o pipefail

# Navigate to script root
parent_path=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)
root_dir=$(realpath "$parent_path/../..")

cd "$root_dir/contracts"

# # Convert hex to 256-bit decimal using Python
DECIMAL_KEY=$(python3 -c "print(int('$OPERATOR_PVT_KEY', 16))")
forge script script/RegisterOperatorToAvs.s.sol:RegisterOperatorToAvs --rpc-url "$RPC_URL" -vvvv --via-ir  --broadcast --sig  "run(uint,uint,uint,uint,uint,uint,uint,uint,uint)" -- "$DECIMAL_KEY" "$G1_X" "$G1_Y" "$G2_X0" "$G2_X1" "$G2_Y0" "$G2_Y1" "$PUBKEY_REGISTRATION_SIGNATURE_X" "$PUBKEY_REGISTRATION_SIGNATURE_Y"
