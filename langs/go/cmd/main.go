package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Layr-Labs/hourglass-monorepo/ponos/pkg/performer/server"
	performerV1 "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/hourglass/v1/performer"
	"go.uber.org/zap"
)

// This offchain binary is run by Operators running the Hourglass Executor. It contains
// the business logic of the AVS and performs worked based on the tasked sent to it.
// The Hourglass Aggregator ingests tasks from the TaskMailbox and distributes work
// to Executors configured to run the AVS Performer. Performers execute the work and
// return the result to the Executor where the result is signed and return to the
// Aggregator to place in the outbox once the signing threshold is met.

type TaskWorker struct {
	logger *zap.Logger
}

func NewTaskWorker(logger *zap.Logger) *TaskWorker {
	return &TaskWorker{
		logger: logger,
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
	// E.g. the Perfomer may validate that the request params are well formed and adhere to a schema.

	return nil
}

func (tw *TaskWorker) HandleTask(t *performerV1.TaskRequest) (*performerV1.TaskResponse, error) {
	tw.logger.Sugar().Infow("Handling task",
		zap.Any("task", t),
	)

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
