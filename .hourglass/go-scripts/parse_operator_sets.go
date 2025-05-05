package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

type OperatorSet struct {
	OperatorSetID int    `toml:"operator_set_id"`
	Description   string `toml:"description"`
	RPCEndpoint   string `toml:"rpc_endpoint"`
	AVS           string `toml:"avs"`
	SubmitWallet  string `toml:"submit_wallet"`
}

type EigenConfig struct {
	OperatorSets map[string]OperatorSet `toml:"operatorsets"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run go-scripts/parse_operator_sets.go <path-to-eigen.toml>")
		os.Exit(1)
	}
	path := os.Args[1]
	var config EigenConfig
	if _, err := toml.DecodeFile(path, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse TOML: %v\n", err)
		os.Exit(1)
	}

	var aliases []string
	for alias := range config.OperatorSets {
		aliases = append(aliases, alias)
	}

	// Output as Makefile-compatible variables
	fmt.Printf("OPERATOR_ALIAS_list := %s\n", strings.Join(aliases, " "))
	for alias, set := range config.OperatorSets {
		prefix := strings.ToUpper(strings.ReplaceAll(alias, "-", "_"))
		fmt.Printf("%s_OPERATOR_SET_ID := %d\n", prefix, set.OperatorSetID)
		fmt.Printf("%s_RPC_ENDPOINT := %s\n", prefix, set.RPCEndpoint)
		fmt.Printf("%s_AVS := %s\n", prefix, set.AVS)
		fmt.Printf("%s_SUBMIT_WALLET := %s\n", prefix, set.SubmitWallet)
	}
}
