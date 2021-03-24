package database

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/memeoAmazonas/demo-2/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func GetDB() (*mongo.Database, error) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "connectiondb",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	uri := common.GetEnv("URL_BD")
	dbName := common.GetEnv("DB_NAME")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		level.Error(logger).Log("connectiondb", err)
		os.Exit(-1)
		return nil, err
	}
	level.Info(logger).Log("message", "Conexion successful")
	return client.Database(dbName), nil
}
