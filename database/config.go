package database

import (
	"context"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	uri := "mongodb://localhost:27017/" //! por alguna razon que desconosco con los test no carga el archivo .env
	dbName := "gokitdbtest"
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
