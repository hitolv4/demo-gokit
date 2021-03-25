package repository

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InvalidMongoId(id string) error {
	if !primitive.IsValidObjectID(id) {
		return errors.New("Invalid id")
	}
	return nil
}
