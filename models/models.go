package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     *string            `bson:"name,omitempty"`
	Phone    *int               `bson:"phone,omitempty"`
	Email    *string            `bson:"email,omitempty"`
	Password *string            `bson:"password,omitempty"`
}
