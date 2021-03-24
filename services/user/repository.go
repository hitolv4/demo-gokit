package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Email    string             `json:"email"`
	Name     string             `json:"name"`
	Password string             `json:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetAll(ctx context.Context) ([]*User, error)
	GetById(ctx context.Context, id string) (User, error)
	DeleteById(ctx context.Context, id string) error
}
