#!/usr/bin/env bash
set -euo pipefail

eigen_toml="eigen.toml"

eval "$(go run ./go-scripts/parse_operator_sets.go "$eigen_toml")"
eval "$(go run ./go-scripts/parse_operator_keys.go "$eigen_toml")"
eval "$(go run ./go-scripts/parse_rpc_url.go "$eigen_toml")"

read -a KEYS_ARRAY <<< "$OPERATOR_KEYS_list"
KEYSTORE_COUNT=${#KEYS_ARRAY[@]}

for (( i=0; i<KEYSTORE_COUNT; i++ )); do
	export OPERATOR_KEY="${KEYS_ARRAY[$i]}"
	echo "🔑 Registering operator $i: $OPERATOR_KEY"

	eval "$(go run ./go-scripts/parse_keystores.go "$eigen_toml")"
	eval "$(RPC_URL=$RPC_URL REGISTRAR_ADDRESS=$REGISTRAR_ADDRESS OPERATOR_KEY=$OPERATOR_KEY go run ./go-scripts/get_pubkey_registration_hash.go)"
	export OPERATOR_INDEX=$i
	eval "$(OPERATOR_BLSPRIVATE_KEY=$OPERATOR_KEY PUBKEY_REGISTRATION_MESSAGE_HASH_X_POINT=$PUBKEY_REGISTRATION_MESSAGE_HASH_X_POINT PUBKEY_REGISTRATION_MESSAGE_HASH_Y_POINT=$PUBKEY_REGISTRATION_MESSAGE_HASH_Y_POINT go run ./go-scripts/sign_pubkey_hash.go)"

	G1_X="$(eval echo \${OPERATOR_${i}_G1_X})"
	G1_Y="$(eval echo \${OPERATOR_${i}_G1_Y})"
	G2_X0="$(eval echo \${OPERATOR_${i}_G2_X0})"
	G2_X1="$(eval echo \${OPERATOR_${i}_G2_X1})"
	G2_Y0="$(eval echo \${OPERATOR_${i}_G2_Y0})"
	G2_Y1="$(eval echo \${OPERATOR_${i}_G2_Y1})"

	export G1_X G1_Y G2_X0 G2_X1 G2_Y0 G2_Y1
	export PUBKEY_REGISTRATION_MESSAGE_HASH_X_POINT PUBKEY_REGISTRATION_MESSAGE_HASH_Y_POINT
	export PUBKEY_REGISTRATION_SIGNATURE_X PUBKEY_REGISTRATION_SIGNATURE_Y

	for alias in $OPERATOR_ALIAS_list; do
		alias_var=$(echo "$alias" | tr a-z- A-Z_)
		operator_set_id=$(eval echo "\${${alias_var}_OPERATOR_SET_ID}")
		OPERATOR_SET_ID=$operator_set_id OPERATOR_PVT_KEY=$OPERATOR_KEY ./.hourglass/scripts/register-operator-to-avs.sh
	done
done
