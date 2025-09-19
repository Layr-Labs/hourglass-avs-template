package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Layr-Labs/hourglass-avs-template/contracts/bindings/l1/helloworldl1"
	"github.com/Layr-Labs/hourglass-avs-template/contracts/bindings/l1/taskavsregistrar"
	"github.com/Layr-Labs/hourglass-monorepo/ponos/pkg/performer/contracts"
	"github.com/Layr-Labs/hourglass-monorepo/ponos/pkg/performer/server"
	performerV1 "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/hourglass/v1/performer"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

// This offchain binary is run by Operators running the Hourglass Executor. It contains
// the business logic of the AVS and performs worked based on the tasked sent to it.
// The Hourglass Aggregator ingests tasks from the TaskMailbox and distributes work
// to Executors configured to run the AVS Performer. Performers execute the work and
// return the result to the Executor where the result is signed and return to the
// Aggregator to place in the outbox once the signing threshold is met.

type TaskWorker struct {
	logger        *zap.Logger
	contractStore *contracts.ContractStore
	l1Client      *ethclient.Client
	l2Client      *ethclient.Client
}

func NewTaskWorker(logger *zap.Logger) *TaskWorker {
	// Initialize contract store from environment variables
	contractStore, err := contracts.NewContractStore()
	if err != nil {
		logger.Warn("Failed to load contract store", zap.Error(err))
	}

	// Initialize Ethereum clients if RPC URLs are provided
	var l1Client, l2Client *ethclient.Client

	if l1RpcUrl := os.Getenv("L1_RPC_URL"); l1RpcUrl != "" {
		l1Client, err = ethclient.Dial(l1RpcUrl)
		if err != nil {
			logger.Error("Failed to connect to L1 RPC", zap.Error(err))
		}
	}

	if l2RpcUrl := os.Getenv("L2_RPC_URL"); l2RpcUrl != "" {
		l2Client, err = ethclient.Dial(l2RpcUrl)
		if err != nil {
			logger.Error("Failed to connect to L2 RPC", zap.Error(err))
		}
	}

	return &TaskWorker{
		logger:        logger,
		contractStore: contractStore,
		l1Client:      l1Client,
		l2Client:      l2Client,
	}
}

func (tw *TaskWorker) ValidateTask(t *performerV1.TaskRequest) error {
	tw.logger.Sugar().Infow("Validating task",
		zap.Any("task", t),
	)

	// ------------------------------------------------------------------------
	// Implement your AVS task validation logic here
	// ------------------------------------------------------------------------
	// This is where the Perfomer will validate the task request data.
	// E.g. the Perfomer may validate that the request params are well-formed and adhere to a schema.

	return nil
}

func (tw *TaskWorker) HandleTask(t *performerV1.TaskRequest) (*performerV1.TaskResponse, error) {
	tw.logger.Sugar().Infow("Handling task",
		zap.Any("task", t),
	)

	// ------------------------------------------------------------------------
	// Example: How to interact with contracts
	// ------------------------------------------------------------------------

	// Example 1: Generate bindings to contracts
	if tw.contractStore != nil {

		taskRegistrarAddr, err := tw.contractStore.GetTaskAVSRegistrar()
		if err != nil {
			tw.logger.Warn("TaskAVSRegistrar not found", zap.Error(err))
		} else {
			tw.logger.Info("TaskAVSRegistrar", zap.String("address", taskRegistrarAddr.Hex()))

			// TaskAVSRegistrar contract binding
			if tw.l1Client != nil {
				registrar, err := taskavsregistrar.NewTaskAVSRegistrar(taskRegistrarAddr, tw.l1Client)
				if err == nil {
					// Call the registrar contract
					_ = registrar
				}
			}
		}

		// Example 2: Get custom contract addresses
		if helloWorldL1, err := tw.contractStore.GetContract("HELLO_WORLD_L1"); err == nil {
			tw.logger.Info("HelloWorldL1 contract", zap.String("address", helloWorldL1.Hex()))

			// Use the address to create a contract binding
			contract, err := helloworldl1.NewHelloWorldL1(helloWorldL1, tw.l1Client)
			if err == nil {
				message, _ := contract.GetMessage(nil)
				tw.logger.Info("Contract message", zap.String("message", message))
			}
		}

		// Example 3: List available contracts
		tw.logger.Info("Available contracts", zap.Strings("contracts", tw.contractStore.ListContracts()))
	}

	// ------------------------------------------------------------------------
	// Implement your AVS logic here
	// ------------------------------------------------------------------------
	// This is where the Performer will do the work and provide compute.
	// E.g. the Perfomer could call an external API, a local service or a script.
	var resultBytes []byte
	return &performerV1.TaskResponse{
		TaskId: t.TaskId,
		Result: resultBytes,
	}, nil
}

func main() {
	ctx := context.Background()
	l, _ := zap.NewProduction()

	w := NewTaskWorker(l)

	pp, err := server.NewPonosPerformerWithRpcServer(&server.PonosPerformerConfig{
		Port:    8080,
		Timeout: 5 * time.Second,
	}, w, l)
	if err != nil {
		panic(fmt.Errorf("failed to create performer: %w", err))
	}

	if err := pp.Start(ctx); err != nil {
		panic(err)
	}
}
