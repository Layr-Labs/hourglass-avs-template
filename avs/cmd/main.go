package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	performerV1 "github.com/Layr-Labs/hourglass-monorepo/ponos/gen/protos/eigenlayer/hourglass/v1/performer"
	"github.com/Layr-Labs/hourglass-monorepo/ponos/pkg/performer/server"
	"github.com/pkg/errors"
	"go.uber.org/zap"
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
	Value *big.Int
}

type TaskResponsePayload struct {
	Result        *big.Int
	UnixTimestamp uint64
}

// parseABIEncodedUint256 parses an ABI-encoded uint256 from the payload
// In ABI encoding, a uint256 is a 32-byte value, padded on the left
func parseABIEncodedUint256(data []byte) (*big.Int, error) {
	if len(data) < 32 {
		return nil, fmt.Errorf("data too short for uint256: %d bytes", len(data))
	}

	// Extract the uint256 value (32 bytes)
	return new(big.Int).SetBytes(data[:32]), nil
}

func (tw *TaskWorker) marshalPayload(t *performerV1.Task) (*TaskRequestPayload, error) {
	if len(t.Payload) == 0 {
		return nil, fmt.Errorf("task payload is empty")
	}

	payloadBytes := t.GetPayload()
	tw.logger.Sugar().Debugw("Raw payload bytes", "bytes", fmt.Sprintf("%x", payloadBytes))

	// Parse the ABI-encoded uint256 value
	value, err := parseABIEncodedUint256(payloadBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode ABI payload")
	}

	return &TaskRequestPayload{
		Value: value,
	}, nil
}

func (tw *TaskWorker) ValidateTask(t *performerV1.Task) error {
	tw.logger.Sugar().Infow("Validating task",
		zap.Any("task", t),
	)
	payload, err := tw.marshalPayload(t)
	if err != nil {
		return errors.Wrap(err, "invalid task payload")
	}

	// Check if Value is nil
	if payload.Value == nil {
		return errors.New("value is nil")
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

	squaredNumber := new(big.Int).Exp(payload.Value, big.NewInt(2), nil)

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
