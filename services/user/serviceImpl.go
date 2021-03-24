package user

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(repository Repository, logger log.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s service) CreateUser(ctx context.Context, name string, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")
	id := primitive.NewObjectID()
	user := User{
		ID:       id,
		Name: name,
		Email:    email,
		Password: password,
	}
	if err := s.repository.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("Create user", id)
	return "Success", nil
}
