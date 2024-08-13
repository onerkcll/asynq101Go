package main

import (
	"asynq101Go/application"
	"asynq101Go/application/core/helpers/logHelpers"
	"asynq101Go/application/core/queueManager"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
	AppConf "github.com/spf13/viper"
)

func createDbCacheTask(key, value string, client *asynq.Client, logger *logrus.Logger) {

	task, err := queueManager.NewDbCacheTask(key, value)
	if err != nil {
		logger.Errorf("could not create task: %v", err)
	}

	info, err := client.Enqueue(task, asynq.Queue("default"), asynq.MaxRetry(0))

	if err != nil {
		logger.Errorf("could not enqueue task: %v", err)
	}
	logger.Infof("[!]PASSED -> enqueued task  ID: %d - queue: %s", info.ID, info.Queue)
}
func main() {
	application.CreateApp(false)
	logger := logHelpers.GetLogger()
	redisAddr := AppConf.GetString("RedisServer")
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()
	key := "key"
	value := "value"
	for i := 0; i < 400; i++ {
		tempKey := fmt.Sprintf("%s%d", key, i)
		tempValue := fmt.Sprintf("%s%d", value, i)
		createDbCacheTask(tempKey, tempValue, client, logger)
	}
}
