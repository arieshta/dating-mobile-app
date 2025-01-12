package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Likes struct {
	ID      primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	UserID  string               `json:"user_id" bson:"user_id"`
	LikeIDs []primitive.ObjectID `json:"like_ids" bson:"like_ids"`
}
