package muxManager

import (
	"asynq101Go/application/core/queueManager"
	"context"
	"github.com/hibiken/asynq"
)

var muxManagerMap = map[string]interface{}{
	queueManager.TypeDbCacheWriter: queueManager.HandleDbCachePayload,
}

func initMux() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	return mux
}
func CreateMuxManager() *asynq.ServeMux {
	mux := initMux()
	for taskType, handler := range muxManagerMap {
		mux.HandleFunc(taskType, handler.(func(context.Context, *asynq.Task) error))
	}
	return mux
}
