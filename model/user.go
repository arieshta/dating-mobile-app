package model

import (
	"time"

	"github.com/arieshta/dating-mobile-app/model/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Email     string             `json:"email" bson:"email"`
	Fullname  string             `json:"fullname" bson:"fullname"`
	Gender    string             `json:"gender" bson:"gender"`
	Picture   string             `json:"picture" bson:"picture"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (user *User) ToProfile() *dto.Profile {
	return &dto.Profile{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Fullname: user.Fullname,
		Gender:   user.Gender,
		Picture:  user.Picture,
		Verified: false,
	}
}
