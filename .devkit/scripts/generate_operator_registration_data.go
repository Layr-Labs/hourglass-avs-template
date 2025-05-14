package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/yaml.v3"
)

// Framework configuration structure
type FrameworkConfig struct {
	TaskSla                 int        `yaml:"taskSla"`
	AggregatorOperatorSetId int        `yaml:"aggregatorOperatorSetId"`
	ExecutorOperatorSetId   int        `yaml:"executorOperatorSetId"`
	Operators               []Operator `yaml:"operators"`
}

type Operator struct {
	Address string `yaml:"address"`
	Socket  string `yaml:"socket"`
}

// DevnetConfig structure
type DevnetConfig struct {
	Context struct {
		ChainID int    `yaml:"chain_id"`
		RpcURL  string `yaml:"rpc_url"`
	} `yaml:"context"`
}

// DeployOutput structure
type DeployOutput struct {
	Addresses struct {
		TaskAVSRegistrar string `json:"taskAVSRegistrar"`
	} `json:"addresses"`
	ChainInfo struct {
		ChainID         int `json:"chainId"`
		DeploymentBlock int `json:"deploymentBlock"`
	} `json:"chainInfo"`
	Parameters struct {
		AllocationManager string `json:"allocationManager"`
		AVS               string `json:"avs"`
	} `json:"parameters"`
}

// Reads the framework YAML file and returns the socket for the given operator
func readFrameworkYaml(operatorAddress common.Address) (string, error) {
	frameworkYaml, err := os.ReadFile("../framework.yaml")
	if err != nil {
		return "", fmt.Errorf("error reading framework.yaml: %v", err)
	}

	var frameworkConfig FrameworkConfig
	err = yaml.Unmarshal(frameworkYaml, &frameworkConfig)
	if err != nil {
		return "", fmt.Errorf("error parsing framework.yaml: %v", err)
	}

	// Find the socket for the operator
	for _, op := range frameworkConfig.Operators {
		opAddr := strings.ToLower(op.Address)
		if strings.HasPrefix(opAddr, "0x") {
			opAddr = opAddr[2:]
		}
		opAddr = "0x" + opAddr

		if strings.EqualFold(opAddr, operatorAddress.Hex()) {
			return op.Socket, nil
		}
	}

	return "", fmt.Errorf("socket not found for operator %s in framework.yaml", operatorAddress.Hex())
}

// Reads the deployment JSON file and returns the TaskAVSRegistrar and AVS addresses
func readDeployOutput() (common.Address, common.Address, error) {
	deployOutputJson, err := os.ReadFile("../../contracts/script/local/output/deploy_avs_l1_output.json")
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("error reading deploy_avs_l1_output.json: %v", err)
	}

	var deployOutput DeployOutput
	err = json.Unmarshal(deployOutputJson, &deployOutput)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("error parsing deploy_avs_l1_output.json: %v", err)
	}

	taskAVSRegistrar := common.HexToAddress(deployOutput.Addresses.TaskAVSRegistrar)
	avsAddress := common.HexToAddress(deployOutput.Parameters.AVS)

	return taskAVSRegistrar, avsAddress, nil
}

// Reads the devnet YAML file and returns the Chain ID and RPC URL
func readDevnetYaml(context string) (*big.Int, string, error) {
	devnetYamlContent, err := os.ReadFile("../../config/contexts/" + context + ".yaml")
	if err != nil {
		return nil, "", fmt.Errorf("error reading devnet.yaml: %v", err)
	}

	// Define a struct that matches the structure of the devnet.yaml file
	// Include only the fields we need to extract
	type DevnetConfig struct {
		Context struct {
			ChainID int    `yaml:"chain_id"`
			RpcURL  string `yaml:"rpc_url"`
		} `yaml:"context"`
	}

	var config DevnetConfig
	err = yaml.Unmarshal(devnetYamlContent, &config)
	if err != nil {
		return nil, "", fmt.Errorf("error parsing devnet.yaml: %v", err)
	}

	// Convert chain_id to big.Int
	chainID := big.NewInt(int64(config.Context.ChainID))
	rpcURL := config.Context.RpcURL

	// Use defaults if values are not present
	if chainID.Cmp(big.NewInt(0)) == 0 {
		chainID = big.NewInt(31337) // Default for local devnet
	}

	if rpcURL == "" {
		rpcURL = "http://localhost:8545" // Default RPC URL
	}

	return chainID, rpcURL, nil
}

// Calls the contract to get the hash point
func getHashPoint(operatorAddress common.Address, rpcURL string, registrarAddress common.Address) (*big.Int, *big.Int, error) {
	// Connect to the Ethereum node
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, nil, fmt.Errorf("error connecting to Ethereum node: %v", err)
	}

	// Define the ABI for calling pubkeyRegistrationMessageHash
	const registrarAbiJSON = `[{"inputs":[{"internalType":"address","name":"operator","type":"address"}],"name":"pubkeyRegistrationMessageHash","outputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
	registrarAbi, err := abi.JSON(strings.NewReader(registrarAbiJSON))
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing ABI: %v", err)
	}

	// Pack the function call data
	callData, err := registrarAbi.Pack("pubkeyRegistrationMessageHash", operatorAddress)
	if err != nil {
		return nil, nil, fmt.Errorf("error packing function call: %v", err)
	}

	// Create the call message
	msg := ethereum.CallMsg{
		To:   &registrarAddress,
		Data: callData,
	}

	// Make the call
	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("error calling contract: %v", err)
	}

	// Unpack the result
	unpacked, err := registrarAbi.Unpack("pubkeyRegistrationMessageHash", result)
	if err != nil {
		return nil, nil, fmt.Errorf("error unpacking result: %v", err)
	}

	// Check if we got the expected two values
	if len(unpacked) != 2 {
		return nil, nil, fmt.Errorf("unexpected result format, expected 2 values, got %d", len(unpacked))
	}

	// Convert the values to big.Int
	hashPointX := unpacked[0].(*big.Int)
	hashPointY := unpacked[1].(*big.Int)

	return hashPointX, hashPointY, nil
}

// Calculate BLS keys and signatures
func generateBLSKeysAndSignature(privateKeyBigInt *big.Int, hashPointX, hashPointY *big.Int) (bn254.G1Affine, bn254.G1Affine, bn254.G2Affine, error) {
	// Convert private key to scalar for BLS operations
	var privateKey fr.Element
	privateKey.SetBigInt(privateKeyBigInt)

	// Create a proper G1 generator point
	var g1Gen bn254.G1Affine
	_, err := g1Gen.X.SetString("1")
	if err != nil {
		return bn254.G1Affine{}, bn254.G1Affine{}, bn254.G2Affine{}, fmt.Errorf("error setting G1 generator X: %v", err)
	}
	_, err = g1Gen.Y.SetString("2")
	if err != nil {
		return bn254.G1Affine{}, bn254.G1Affine{}, bn254.G2Affine{}, fmt.Errorf("error setting G1 generator Y: %v", err)
	}

	// Create a proper G2 generator point
	var g2Gen bn254.G2Affine
	// These are the coordinates of the standard BN254 G2 generator
	_, err = g2Gen.X.A0.SetString("10857046999023057135944570762232829481370756359578518086990519993285655852781")
	if err != nil {
		return bn254.G1Affine{}, bn254.G1Affine{}, bn254.G2Affine{}, fmt.Errorf("error setting G2 generator X.A0: %v", err)
	}
	_, err = g2Gen.X.A1.SetString("11559732032986387107991004021392285783925812861821192530917403151452391805634")
	if err != nil {
		return bn254.G1Affine{}, bn254.G1Affine{}, bn254.G2Affine{}, fmt.Errorf("error setting G2 generator X.A1: %v", err)
	}
	_, err = g2Gen.Y.A0.SetString("8495653923123431417604973247489272438418190587263600148770280649306958101930")
	if err != nil {
		return bn254.G1Affine{}, bn254.G1Affine{}, bn254.G2Affine{}, fmt.Errorf("error setting G2 generator Y.A0: %v", err)
	}
	_, err = g2Gen.Y.A1.SetString("4082367875863433681332203403145435568316851327593401208105741076214120093531")
	if err != nil {
		return bn254.G1Affine{}, bn254.G1Affine{}, bn254.G2Affine{}, fmt.Errorf("error setting G2 generator Y.A1: %v", err)
	}

	// Calculate public keys
	// G1 public key
	var pubkeyG1 bn254.G1Affine
	pubkeyG1.ScalarMultiplication(&g1Gen, privateKeyBigInt)

	// G2 public key
	var pubkeyG2 bn254.G2Affine
	pubkeyG2.ScalarMultiplication(&g2Gen, privateKeyBigInt)

	// Use the hash point coordinates directly from the contract
	var hashPoint bn254.G1Affine

	// Set the point coordinates directly using SetString
	_, err = hashPoint.X.SetString(hashPointX.String())
	if err != nil {
		return bn254.G1Affine{}, bn254.G1Affine{}, bn254.G2Affine{}, fmt.Errorf("error setting hash point X: %v", err)
	}

	_, err = hashPoint.Y.SetString(hashPointY.String())
	if err != nil {
		return bn254.G1Affine{}, bn254.G1Affine{}, bn254.G2Affine{}, fmt.Errorf("error setting hash point Y: %v", err)
	}

	// Sign the message (scalar multiplication of the hash point by private key)
	var signature bn254.G1Affine
	signature.ScalarMultiplication(&hashPoint, privateKeyBigInt)

	return signature, pubkeyG1, pubkeyG2, nil
}

// Format registration parameters
func formatRegistrationParams(operatorSocket string, signature, pubkeyG1 bn254.G1Affine, pubkeyG2 bn254.G2Affine) (string, string) {
	// Format the BLS parameters for Solidity
	pubkeyRegistrationParams := formatForSolidity(signature, pubkeyG1, pubkeyG2)

	// Format the operator registration params (encoded socket + pubkey registration params)
	operatorRegistrationParams := fmt.Sprintf(
		"0x%s%s",
		encodeString(operatorSocket),
		strings.TrimPrefix(pubkeyRegistrationParams, "0x"),
	)

	// Format the final data field for RegisterParams
	registerParamsData := fmt.Sprintf("0x%064x%s", len(operatorRegistrationParams)/2-1, strings.TrimPrefix(operatorRegistrationParams, "0x"))

	return pubkeyRegistrationParams, registerParamsData
}

// Format the BLS parameters for Solidity
func formatForSolidity(signature, pubkeyG1 bn254.G1Affine, pubkeyG2 bn254.G2Affine) string {
	// Get G1 point coordinates
	sigX, sigY := signature.X, signature.Y
	pkG1X, pkG1Y := pubkeyG1.X, pubkeyG1.Y

	// Get G2 point coordinates
	// Note: G2 points use two coordinates for each of X and Y
	pkG2X0, pkG2X1 := pubkeyG2.X.A0, pubkeyG2.X.A1
	pkG2Y0, pkG2Y1 := pubkeyG2.Y.A0, pubkeyG2.Y.A1

	// Convert to bytes and format as hex strings
	params := []string{
		// Signature (G1 point)
		formatBigInt(sigX.BigInt(new(big.Int))),
		formatBigInt(sigY.BigInt(new(big.Int))),

		// PubkeyG1 (G1 point)
		formatBigInt(pkG1X.BigInt(new(big.Int))),
		formatBigInt(pkG1Y.BigInt(new(big.Int))),

		// PubkeyG2 (G2 point)
		// Note: Order matters due to the way EVM expects G2 points
		formatBigInt(pkG2X1.BigInt(new(big.Int))), // X.A1 first
		formatBigInt(pkG2X0.BigInt(new(big.Int))), // X.A0 second
		formatBigInt(pkG2Y1.BigInt(new(big.Int))), // Y.A1 first
		formatBigInt(pkG2Y0.BigInt(new(big.Int))), // Y.A0 second
	}

	// Join all hex strings without separators, add 0x prefix
	return "0x" + strings.Join(params, "")
}

// Format a big.Int as a 32-byte hex string without 0x prefix
func formatBigInt(b *big.Int) string {
	// Ensure 32 bytes (64 hex chars)
	return fmt.Sprintf("%064s", hex.EncodeToString(b.Bytes()))
}

// Encode string to hex with proper ABI encoding
func encodeString(s string) string {
	bytes := []byte(s)
	hexStr := hex.EncodeToString(bytes)

	// Calculate and pad the length
	length := len(bytes)
	lengthHex := fmt.Sprintf("%064x", length)

	// Return the hex-encoded string with the length prefix
	return lengthHex + hexStr
}

func main() {
	if len(os.Args) < 4 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Usage: go run generate_operator_registration_data.go <context> <operator_address> <bls_private_key>")
		fmt.Println("The BLS private key should be provided as a decimal string.")
		fmt.Println("\nExample:")
		fmt.Println("  go run generate_operator_registration_data.go devnet 0x90F79bf6EB2c4f870365E785982E1f101E93b906 1986833214941433436267332083732250924404411626145227166241876916004066224403")
		os.Exit(1)
	}

	context := os.Args[1]

	operatorAddress := common.HexToAddress(os.Args[2])
	privateKeyBigInt, ok := new(big.Int).SetString(os.Args[3], 10)
	if !ok {
		fmt.Println("Invalid BLS private key")
		os.Exit(1)
	}

	// Read chain ID and RPC URL from devnet config
	chainID, rpcURL, err := readDevnetYaml(context)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	// Read operator socket from framework.yaml
	operatorSocket, err := readFrameworkYaml(operatorAddress)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	// Read registrar and AVS addresses from deployment output JSON
	taskAVSRegistrar, _, err := readDeployOutput()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Using Chain ID: %s and RPC URL: %s\n", chainID.String(), rpcURL)

	// Get hash point from contract
	hashPointX, hashPointY, err := getHashPoint(operatorAddress, rpcURL, taskAVSRegistrar)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Hash point: X=[%s], Y=[%s]\n", hashPointX.String(), hashPointY.String())

	// Generate BLS keys and signature
	signature, pubkeyG1, pubkeyG2, err := generateBLSKeysAndSignature(privateKeyBigInt, hashPointX, hashPointY)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	// Format registration parameters
	pubkeyRegistrationParams, registerParamsData := formatRegistrationParams(operatorSocket, signature, pubkeyG1, pubkeyG2)

	fmt.Println("CONTEXT:", context)
	fmt.Println("TASK_AVS_REGISTRAR:", taskAVSRegistrar.Hex())
	fmt.Println("OPERATOR:", operatorAddress.Hex())
	fmt.Println("SOCKET:", operatorSocket)
	fmt.Println("PUBKEY_REGISTRATION_PARAMS:", pubkeyRegistrationParams)
	fmt.Println("\nREGISTER_PARAMS_DATA:", registerParamsData)
}
