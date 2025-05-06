package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type OperatorAllocations struct {
	Strategies    []string `toml:"strategies"`
	TaskExecutors []string `toml:"task-executors"`
	Aggregators   []string `toml:"aggregators"`
}

type OperatorConfig struct {
	Keys       []string             `toml:"keys"`
	Allocations OperatorAllocations `toml:"allocations"`
}

type OperatorSetConfig struct {
	OperatorSetID int    `toml:"operator_set_id"`
	AVS           string `toml:"avs"`
}

type EigenConfig struct {
	Operator      OperatorConfig               `toml:"operator"`
	OperatorSets  map[string]OperatorSetConfig `toml:"operatorsets"`
	AliasMapping  map[string]string            `toml:"operatorset_aliases"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run parse_operator_allocations.go <eigen.toml>")
		os.Exit(1)
	}
	tomlPath := os.Args[1]

	var cfg EigenConfig
	if _, err := toml.DecodeFile(tomlPath, &cfg); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to decode TOML: %v\n", err)
		os.Exit(1)
	}

	strategies := cfg.Operator.Allocations.Strategies
	fmt.Printf("export STRATEGY_COUNT=%d\n", len(strategies))

	for stratIndex, strategy := range strategies {
		fmt.Printf("export STRATEGY_%d_ADDRESS=%s\n", stratIndex, strategy)
		groupIndex := 0
		for alias, field := range cfg.AliasMapping {
			allocs, ok := getAllocationsByField(cfg.Operator.Allocations, field)
			if !ok || stratIndex >= len(allocs) {
				continue
			}
			setInfo, exists := cfg.OperatorSets[field]
			if !exists {
				fmt.Fprintf(os.Stderr, "Missing operator set for alias: %s (field: %s)\n", alias, field)
				continue
			}

			prefix := fmt.Sprintf("STRATEGY_%d_GROUP_%d", stratIndex, groupIndex)
			fmt.Printf("export %s_ALLOCATION=%s\n", prefix, allocs[stratIndex])
			fmt.Printf("export %s_OPERATOR_SET_ID=%d\n", prefix, setInfo.OperatorSetID)
			fmt.Printf("export %s_AVS_ADDRESS=%s\n", prefix, setInfo.AVS)

			if stratIndex < len(cfg.Operator.Keys) {
				fmt.Printf("export %s_OPERATOR_PVT_KEY=%s\n", prefix, cfg.Operator.Keys[stratIndex])
			}
			groupIndex++
		}
		fmt.Printf("export STRATEGY_%d_ALLOCATION_GROUP_COUNT=%d\n", stratIndex, groupIndex)
	}
}

func getAllocationsByField(allocs OperatorAllocations, field string) ([]string, bool) {
	switch field {
	case "task-executors":
		return allocs.TaskExecutors, true
	case "aggregators":
		return allocs.Aggregators, true
	default:
		return nil, false
	}
}
