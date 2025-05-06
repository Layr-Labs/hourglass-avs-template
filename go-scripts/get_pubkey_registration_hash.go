package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)


func main() {
	rpcURL := os.Getenv("RPC_URL")
registrarAddr := os.Getenv("REGISTRAR_ADDRESS")
operatorKey := os.Getenv("OPERATOR_KEY")

operatorAddr, err := getOperatorAddressFromHexKey(operatorKey)

x, y, err := getPubkeyRegistrationMessageHash(rpcURL, common.HexToAddress(registrarAddr), operatorAddr)
if err != nil {
	log.Fatalf("Failed to get pubkeyRegistrationMessageHash for operator %v", err)
}
fmt.Printf("export PUBKEY_REGISTRATION_MESSAGE_HASH_X_POINT=%s\n", x.String())
fmt.Printf("export PUBKEY_REGISTRATION_MESSAGE_HASH_Y_POINT=%s\n", y.String())

	
}


func getPubkeyRegistrationMessageHash(rpcURL string, registrarAddr common.Address, operatorAddr common.Address) (*big.Int, *big.Int, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RPC: %w", err)
	}
	defer client.Close()

	const abiJSON = `
	[
	  {
		"inputs": [{"internalType": "address", "name": "operator", "type": "address"}],
		"name": "pubkeyRegistrationMessageHash",
		"outputs": [
		  {"internalType": "uint256", "name": "x", "type": "uint256"},
		  {"internalType": "uint256", "name": "y", "type": "uint256"}
		],
		"stateMutability": "view",
		"type": "function"
	  }
	]`

	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	input, err := parsedABI.Pack("pubkeyRegistrationMessageHash", operatorAddr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to pack input: %w", err)
	}

	callMsg := ethereum.CallMsg{
		To:   &registrarAddr,
		Data: input,
	}

	out, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("contract call failed: %w", err)
	}

	var result struct {
		X *big.Int
		Y *big.Int
	}
	if err := parsedABI.UnpackIntoInterface(&result, "pubkeyRegistrationMessageHash", out); err != nil {
		return nil, nil, fmt.Errorf("unpack failed: %w", err)
	}

	return result.X, result.Y, nil
}


func getOperatorAddressFromHexKey(hexKey string) (common.Address, error) {
	privKey, err := crypto.HexToECDSA(strings.TrimPrefix(hexKey, "0x"))
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to parse private key: %w", err)
	}
	return crypto.PubkeyToAddress(privKey.PublicKey), nil
}