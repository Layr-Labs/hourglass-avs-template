package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/BurntSushi/toml"
)

type OperatorConfig struct {
	Keys []string `toml:"keys"`
}

type EigenConfig struct {
	Operator OperatorConfig `toml:"operator"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run parse_operator_keys.go <path-to-eigen.toml>")
		os.Exit(1)
	}
	path := os.Args[1]

	var config EigenConfig
	if _, err := toml.DecodeFile(path, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse TOML: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("export OPERATOR_KEYS_list=\"%s\"\n", strings.Join(config.Operator.Keys, " "))
}
