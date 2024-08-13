package application

import (
	"asynq101Go/application/core/helpers/logHelpers"
	"asynq101Go/application/core/muxManager"
	"github.com/hibiken/asynq"
	AppConf "github.com/spf13/viper"
	"log"
)

func CreateApp(createServer bool) {
	// load config
	AppConf.AddConfigPath("config")
	AppConf.SetConfigName("dev_conf")
	AppConf.SetConfigName("prp_conf")
	_ = AppConf.ReadInConfig()
	_ = AppConf.MergeInConfig()
	// Initialize Logging
	logHelpers.InitializeLogger()
	if createServer {
		logger := logHelpers.GetLogger()
		redisSrv := AppConf.GetString("RedisServer")
		logger.Infof("Redis Server: %s", redisSrv)
		srv := asynq.NewServer(
			asynq.RedisClientOpt{Addr: redisSrv},
			asynq.Config{
				Concurrency: 1,
				Queues: map[string]int{
					"default": 1,
				},
			},
		)
		mux := muxManager.CreateMuxManager()

		if err := srv.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	}

}
