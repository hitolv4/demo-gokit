package user

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoint struct {
	CreateUser endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoint {
	return Endpoint{
		CreateUser: makeCreateUser(s),
	}
}

func makeCreateUser(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserRequest)
		ok, err := s.CreateUser(ctx, req.Name, req.Email, req.Password)
		return UserResponse{
			OK: ok,
		}, err
	}
}
