package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Item represents an item with name, token, and URL
type Cluster struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Token string             `bson:"token" json:"token"`
	URL   string             `bson:"url" json:"url"`
}
