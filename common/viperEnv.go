package common

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/spf13/viper"
	"os"
)

func GetEnv(key string) string {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "common",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		level.Error(logger).Log("getEnv", err)
		os.Exit(-1)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		level.Error(logger).Log("message", "invalid key")
	}
	return value
}
