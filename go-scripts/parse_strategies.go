package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type OperatorAllocations struct {
	Strategies []string `toml:"strategies"`
}

type OperatorConfig struct {
	Allocations OperatorAllocations `toml:"allocations"`
}

type EigenConfig struct {
	Operator OperatorConfig `toml:"operator"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run parse_strategies.go <eigen.toml>")
		os.Exit(1)
	}
	tomlPath := os.Args[1]

	var cfg EigenConfig
	if _, err := toml.DecodeFile(tomlPath, &cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to decode TOML: %v\n", err)
		os.Exit(1)
	}

	strategies := cfg.Operator.Allocations.Strategies

	fmt.Printf("export STRATEGY_COUNT=%d\n", len(strategies))
	for i, strategy := range strategies {
		fmt.Printf("export STRATEGY_%d_ADDRESS=%s\n", i, strategy)
	}
}
