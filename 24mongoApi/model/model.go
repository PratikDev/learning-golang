package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Netflix struct {
	ID         bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MovieTitle string        `json:"title,omitempty"`
	Watched    bool          `json:"watched,omitempty"`
}
