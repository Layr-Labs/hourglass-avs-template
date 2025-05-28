package main

import (
	"testing"

	performerV1 "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/hourglass/v1/performer"
	"go.uber.org/zap"
)

func Test_TaskRequestPayload(t *testing.T) {
	// ------------------------------------------------------------------------
	// Write your test cases here
	// ------------------------------------------------------------------------

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf("Failed to create logger: %v", err)
	}

	taskWorker := NewTaskWorker(logger)

	taskRequest := &performerV1.TaskRequest{
		TaskId:   []byte("test-task-id"),
		Payload:  []byte("test-data"),
		Metadata: []byte("test-metadata"),
	}

	err = taskWorker.ValidateTask(taskRequest)
	if err != nil {
		t.Errorf("ValidateTask failed: %v", err)
	}

	resp, err := taskWorker.HandleTask(taskRequest)
	if err != nil {
		t.Errorf("HandleTask failed: %v", err)
	}

	t.Logf("Response: %v", resp)
}
