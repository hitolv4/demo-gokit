package user

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/mongo"
)

var repoError = errors.New("unable to handle repo request")

type repository struct {
	db     *mongo.Collection
	logger log.Logger
}

func NewRepository(db *mongo.Collection, logger log.Logger) Repository {
	return &repository{
		db:     db,
		logger: log.With(logger, "repository", "sql"),
	}
}
func (repository *repository) CreateUser(ctx context.Context, user User) error {
	if user.Name == "" || user.Password == "" || user.Email == "" {
		return repoError
	}

	_, err := repository.db.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
