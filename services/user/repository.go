package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty"`
	Email    string             `json:"email"`
	Name     string             `json:"name"`
	Password string             `json:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
}
