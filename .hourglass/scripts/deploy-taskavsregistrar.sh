set -o errexit -o nounset -o pipefail

# Navigate to script root
parent_path=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)
root_dir=$(realpath "$parent_path/../..")

cd "$root_dir/contracts"

forge script script/DeployTaskAVSRegistrar.s.sol:DeployTaskAVSRegistrar --rpc-url "$RPC_URL" -vvvv  --broadcast --sig "run()" 
