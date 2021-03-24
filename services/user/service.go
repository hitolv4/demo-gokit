package user

import "context"

type Service interface {
	CreateUser(ctx context.Context, name string, email string, password string) (string, error)
}
