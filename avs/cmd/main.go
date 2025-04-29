// This main.go file exists to show you which components need to be implemented
// to create a valid Ponos Performer so that Operators can run your AVS software.
//
// Think of this as a template to get started; you are more than welcome to
// structure your project as you see fit.
package main

import (
	"context"
	"fmt"
	performerV1 "github.com/Layr-Labs/hourglass-monorepo/ponos/gen/protos/eigenlayer/hourglass/v1/performer"
	"github.com/Layr-Labs/hourglass-monorepo/ponos/pkg/performer/server"
	"go.uber.org/zap"
	"time"
)

// TaskWorker is a struct that implements the TaskWorker interface
type TaskWorker struct {
	logger *zap.Logger
}

func NewTaskWorker(logger *zap.Logger) *TaskWorker {
	return &TaskWorker{
		logger: logger,
	}
}

// ValidateTask validates the task payload and returns an error
// if it is invalid
func (tw *TaskWorker) ValidateTask(t *performerV1.Task) error {
	// ------------------------------------------------------------------------
	// Add your AVS task validation logic here.
	// ------------------------------------------------------------------------
	return nil
}

// HandleTask handles the task and returns the result
func (tw *TaskWorker) HandleTask(t *performerV1.Task) (*performerV1.TaskResult, error) {
	// ------------------------------------------------------------------------
	// Add your AVS task handling logic here
	// ------------------------------------------------------------------------

	var responseBytes []byte

	return &performerV1.TaskResult{
		TaskId: t.TaskId,
		Result: responseBytes,
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
