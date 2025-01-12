package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	SignUpBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
		Gender   string `json:"gender"`
		Picture  string `json:"picture"`
	}

	LoginBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Profile struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Fullname string `json:"fullname"`
		Gender   string `json:"gender"`
		Picture  string `json:"picture"`
		Verified bool   `json:"verified"`
	}

	GetRandomFilter struct {
		Gender      string `bson:"gender"`
		ExcludedIds []primitive.ObjectID
	}
)
