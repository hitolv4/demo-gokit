package user

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	if er := userIsValid(user); er != nil {
		return er
	}
	_, err := repository.db.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (repository *repository) GetAll(ctx context.Context) ([]*User, error) {
	var users []*User
	current, err := repository.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, repoError
	}
	for current.Next(ctx) {
		var user User
		if err := current.Decode(&user); err != nil {
			return users, err
		}
		users = append(users, &user)
	}
	if err := current.Err(); err != nil {
		return users, err
	}
	_ = current.Close(ctx)
	if len(users) == 0 {
		return users, mongo.ErrNoDocuments
	}
	return users, nil
}
func (repository *repository) GetById(ctx context.Context, id string) (User, error) {

	user := User{}
	if err := invalidId(id); err != nil {
		return user, err
	}
	ID, _ := primitive.ObjectIDFromHex(id)
	if err := repository.db.FindOne(ctx, bson.M{"_id": ID}).Decode(&user); err != nil {
		return user, err
	}
	return user, nil

}

func userIsValid(user User) error {
	if user.Name == "" {
		return errors.New("Name is mandatory")
	}

	if user.Password == "" {
		return errors.New("Password is mandatory")
	}
	if user.Email == "" {
		return errors.New("Email is mandatory")
	}
	return nil
}

func invalidId(id string) error {
	if !primitive.IsValidObjectID(id) {
		return errors.New("Invalid id")
	}
	return nil
}

func (repository *repository) DeleteById(ctx context.Context, id string) error {
	if err := invalidId(id); err != nil {
		return err
	}
	return nil
}
