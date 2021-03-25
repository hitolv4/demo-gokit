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

func (s service) GetById(ctx context.Context, id string) (User, error) {
	logger := log.With(s.logger, "method", "GetById")

	user, err := s.repository.GetById(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return User{}, err
	}
	logger.Log("Get", "Success")
	return user, nil
}

func (s service) CreateUser(ctx context.Context, name string, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")
	user := User{
		ID:       primitive.NewObjectID(),
		Name:     name,
		Email:    email,
		Password: password,
	}
	if err := s.repository.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("Create", "Success")
	return "user created successfully", nil
}

func (s service) Delete(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "Delete")
	if err := s.repository.DeleteById(ctx, id); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("Delete", "Success")
	return "user delete successfully", nil
}
func (s service) GetAll(ctx context.Context) ([]User, error) {
	logger := log.With(s.logger, "method", "GetAll")
	users, err := s.repository.GetAll(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	logger.Log("GetAll", "Success")
	return users, nil
}
