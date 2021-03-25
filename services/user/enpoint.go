package user

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoint struct {
	CreateUser  endpoint.Endpoint
	Delete      endpoint.Endpoint
	GetAllUser  endpoint.Endpoint
	GetUserById endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoint {
	return Endpoint{
		CreateUser:  makeCreateUser(s),
		Delete:      makeDeleteUser(s),
		GetAllUser:  makeGetAll(s),
		GetUserById: makeGetUserById(s),
	}
}

func makeCreateUser(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserRequest)
		ok, err := s.CreateUser(ctx, req.Name, req.Email, req.Password)
		return UserResponse{
			Message: ok,
		}, err
	}
}

func makeGetUserById(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(IdRequest)
		user, err := s.GetById(ctx, req.Id)
		return UserResponse{
			Message: "user get success",
			User:    &user,
		}, err
	}
}
func makeDeleteUser(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(IdRequest)
		ok, err := s.Delete(ctx, req.Id)
		return UserResponse{
			Message: ok,
		}, err
	}
}

func makeGetAll(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.GetAll(ctx)
		return UserResponse{
			Message:  "All user success",
			UserList: users,
		}, err
	}
}
