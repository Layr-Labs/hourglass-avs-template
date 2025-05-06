package main

import (
	"context"
	"encoding/json"
	"fmt"
	performerV1 "github.com/Layr-Labs/hourglass-monorepo/ponos/gen/protos/eigenlayer/hourglass/v1/performer"
	"github.com/Layr-Labs/hourglass-monorepo/ponos/pkg/performer/server"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/big"
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

type TaskRequestPayload struct {
	NumberToBeSquared json.Number `json:"numberToBeSquared"`
}

func (trp *TaskRequestPayload) GetBigInt() (*big.Int, error) {
	i, success := new(big.Int).SetString(trp.NumberToBeSquared.String(), 10)
	if !success {
		return nil, fmt.Errorf("failed to convert json.Number to big.Int")
	}
	return i, nil
}

type TaskResponsePayload struct {
	Result        *big.Int
	UnixTimestamp uint64
}

func (tw *TaskWorker) marshalPayload(t *performerV1.Task) (*TaskRequestPayload, error) {
	if len(t.Payload) == 0 {
		return nil, fmt.Errorf("task payload is empty")
	}

	payloadBytes := t.GetPayload()
	var payload *TaskRequestPayload
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal payload")
	}
	return payload, nil
}

func (tw *TaskWorker) ValidateTask(t *performerV1.Task) error {
	tw.logger.Sugar().Infow("Validating task",
		zap.Any("task", t),
	)
	payload, err := tw.marshalPayload(t)
	if err != nil {
		return errors.Wrap(err, "invalid task payload")
	}

	_, err = payload.GetBigInt()
	if err != nil {
		return errors.Wrap(err, "failed to get big.Int from payload")
	}

	return nil
}

func (tw *TaskWorker) HandleTask(t *performerV1.Task) (*performerV1.TaskResult, error) {
	tw.logger.Sugar().Infow("Handling task",
		zap.Any("task", t),
	)
	payload, err := tw.marshalPayload(t)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal payload")
	}

	i, err := payload.GetBigInt()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get big.Int from payload")
	}

	squaredNumber := new(big.Int).Exp(i, big.NewInt(2), nil)

	responsePayload := &TaskResponsePayload{
		Result:        squaredNumber,
		UnixTimestamp: uint64(time.Now().Unix()),
	}
	responseBytes, err := json.Marshal(responsePayload)

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
