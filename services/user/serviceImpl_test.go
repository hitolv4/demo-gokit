package user

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/memeoAmazonas/demo-2/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreateUser(t *testing.T) {
	user := map[string]string{
		"name":     "test service",
		"email":    "test@test.com",
		"password": "password",
	}
	t.Run("Passing all params", func(t *testing.T) {
		srv, ctx := Setup()
		got, _ := srv.CreateUser(ctx, user["name"], user["email"], user["password"])
		want := "user create success"
		assertString(t, got, want)
	})
	t.Run("missing name", func(t *testing.T) {
		srv, ctx := Setup()
		_, err := srv.CreateUser(ctx, "", user["email"], user["password"])
		want := "Name is mandatory"
		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertString(t, err.Error(), want)
	})
	t.Run("missing email", func(t *testing.T) {
		srv, ctx := Setup()
		_, err := srv.CreateUser(ctx, user["name"], "", user["password"])
		want := "Email is mandatory"
		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertString(t, err.Error(), want)
	})
	t.Run("missing password", func(t *testing.T) {
		srv, ctx := Setup()
		_, err := srv.CreateUser(ctx, user["name"], user["email"], "")
		want := "Password is mandatory"
		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertString(t, err.Error(), want)
	})
}

func TestGETUsers(t *testing.T)   {}
func TestGetUser(t *testing.T)    {}
func TestUpdateUser(t *testing.T) {}
func TestDeleteUser(t *testing.T) {}

func assertString(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got: %s want: %s", got, want)
	}
}

func Setup() (srv Service, ctx context.Context) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "test service",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("message", "test service started")
	defer level.Info(logger).Log("message", "test service end")
	var db *mongo.Database
	{
		var err error
		db, err = database.GetDB()
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

	}
	userCollection := db.Collection("user")
	repository := NewRepository(userCollection, logger)

	return NewService(repository, logger), context.Background()
}

//*IMPORTANTE NO CORRER LA PRUEBA SI EL SERVIDOR ESTA FUNCIONANDO
