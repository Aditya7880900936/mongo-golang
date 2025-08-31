package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name  string             `json:"name"`
    Email string             `json:"email"`
}
