package user

import "context"

type Service interface {
	CreateUser(ctx context.Context, name string, email string, password string) (string, error)
	GetById(ctx context.Context, id string) (User, error)
	Delete(ctx context.Context, id string) (string, error)
	GetAll(ctx context.Context) ([]User, error)
}
