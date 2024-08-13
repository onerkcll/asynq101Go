package queueManager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
)

const (
	TypeDbCacheWriter = "key:value"
)

type DbCachePayload struct {
	Key   string
	Value string
}

// NewDbCacheTask creates a new task to write a key-value pair to the database cache.
func NewDbCacheTask(key string, value string) (*asynq.Task, error) {
	payload, err := json.Marshal(DbCachePayload{Key: key, Value: value})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeDbCacheWriter, payload), nil
}

// HandleDbCachePayload writes a key-value pair to the database cache.
func HandleDbCachePayload(ctx context.Context, t *asynq.Task) error {
	var p DbCachePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Received Key: %s, Received Value: %s", p.Key, p.Value)
	// Email delivery code ...
	return nil
}
