package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/memeoAmazonas/demo-2/database"
	"github.com/memeoAmazonas/demo-2/services/user"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var httpAddr = flag.String("http", ":8888", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "user",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("message", "service started")
	defer level.Info(logger).Log("message", "service end")
	ctx := context.Background()
	var db *mongo.Database
	{
		var err error
		db, err = database.GetDB()
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}
	//defer db.Disconnect(ctx)

	flag.Parse()

	var srv user.Service
	{
		userCollection := db.Collection("user")
		repository := user.NewRepository(userCollection, logger)
		srv = user.NewService(repository, logger)
	}
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := user.MakeEndpoints(srv)
	go func() {
		fmt.Println("listen on port", *httpAddr)
		handler := user.NewHttpServer(ctx,endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()
	level.Error(logger).Log("exit", <-errs)
}
