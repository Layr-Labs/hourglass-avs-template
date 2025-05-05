package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type EnvConfig struct {
	ChainArgs  []string `toml:"chain_args"`
	ChainImage string   `toml:"chain_image"`
}

type EigenConfig struct {
	Env map[string]EnvConfig `toml:"env"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: parse_rpc_url.go <path-to-eigen.toml>")
		os.Exit(1)
	}
	var cfg EigenConfig
	if _, err := toml.DecodeFile(os.Args[1], &cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse TOML: %v\n", err)
		os.Exit(1)
	}

	// Default to Anvil RPC port
	fmt.Printf("RPC_URL=http://localhost:8545\n")
}