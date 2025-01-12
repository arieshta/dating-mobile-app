package model

type Premium struct {
	UserId  string `bson:"user_id"`
	Premium bool   `bson:"premium"`
}
