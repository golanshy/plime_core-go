package id_dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Id struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"_id, omitempty"`
}
