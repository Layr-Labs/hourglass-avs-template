package main

import (
	"context"
	"fmt"
	performerV1 "github.com/Layr-Labs/hourglass-monorepo/ponos/gen/protos/eigenlayer/hourglass/v1/performer"
	"github.com/Layr-Labs/hourglass-monorepo/ponos/pkg/performer/server"
	"go.uber.org/zap"
	"time"
)

type TaskWorker struct {
	logger *zap.Logger
}

func NewTaskWorker(logger *zap.Logger) *TaskWorker {
	return &TaskWorker{
		logger: logger,
	}
}

func (tw *TaskWorker) ValidateTask(t *performerV1.Task) error {
	tw.logger.Sugar().Infow("Validating task",
		zap.Any("task", t),
	)
	// ------------------------------------------------------------------------
	// Implement your AVS task validation logic here
	// ------------------------------------------------------------------------

	return nil
}

func (tw *TaskWorker) HandleTask(t *performerV1.Task) (*performerV1.TaskResult, error) {
	tw.logger.Sugar().Infow("Handling task",
		zap.Any("task", t),
	)

	// ------------------------------------------------------------------------
	// Implement your AVS logic here
	// ------------------------------------------------------------------------
	var resultBytes []byte

	return &performerV1.TaskResult{
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
