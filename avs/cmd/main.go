package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Layr-Labs/hourglass-monorepo/ponos-performer/go/pkg/server"
	"github.com/Layr-Labs/hourglass-monorepo/ponos/pkg/performer"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

type TaskWorker struct{}

type TaskRequestPayload struct {
	Message string
}

type TaskResponsePayload struct {
	Message       string
	UnixTimestamp uint64
}

func (tw *TaskWorker) marshalPayload(t *performer.Task) (*TaskRequestPayload, error) {
	if len(t.Payload) == 0 {
		return nil, fmt.Errorf("task payload is empty")
	}

	payloadBytes, err := t.GetPayloadBytes()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get payload bytes")
	}
	var payload *TaskRequestPayload
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal payload")
	}
	return payload, nil
}

func (tw *TaskWorker) ValidateTask(t *performer.Task) error {
	if _, err := tw.marshalPayload(t); err != nil {
		return errors.Wrap(err, "invalid task payload")
	}
	return nil
}

func (tw *TaskWorker) HandleTask(t *performer.Task) (*performer.TaskResult, error) {
	payload, err := tw.marshalPayload(t)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal payload")
	}

	responsePayload := &TaskResponsePayload{
		Message:       fmt.Sprintf("Hello %s", payload.Message),
		UnixTimestamp: uint64(time.Now().Unix()),
	}
	responseBytes, err := json.Marshal(responsePayload)

	return performer.NewTaskResult(t.TaskID, t.Avs, t.OperatorSetID, responseBytes), nil
}

func main() {
	ctx := context.Background()
	l, _ := zap.NewProduction()

	w := &TaskWorker{}

	pp := server.NewPonosPerformer(&server.PonosPerformerConfig{
		Port:    8080,
		Timeout: 5 * time.Second,
	}, w, l)

	if err := pp.StartHttpServer(ctx); err != nil {
		panic(err)
	}
}
