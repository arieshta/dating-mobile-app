package users

import (
	"github.com/arieshta/dating-mobile-app/internal/services/premium"
	"github.com/arieshta/dating-mobile-app/internal/services/users"
	"github.com/arieshta/dating-mobile-app/model"
)

type (
	Controller struct {
		sharedHolder *model.SharedHolder

		userService users.Service
		premiumService premium.Service
	}
)

func NewController(holder *model.SharedHolder) *Controller {
	UserService, err := users.NewService(holder)
	if err != nil {
		panic(err)
	}

	PremiumService, err := premium.NewService(holder)
	if err != nil {
		panic(err)
	}

	return &Controller{
		sharedHolder: holder,
		userService:  UserService,
		premiumService:  PremiumService,
	}
}