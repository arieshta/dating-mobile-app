package premium

import (
	"github.com/arieshta/dating-mobile-app/internal/services/premium"
	"github.com/arieshta/dating-mobile-app/model"
)

type (
	Controller struct {
		sharedHolder *model.SharedHolder

		premiumService premium.Service
	}
)

func NewController(holder *model.SharedHolder) *Controller {
	PremiumService, err := premium.NewService(holder)
	if err != nil {
		panic(err)
	}

	return &Controller{
		sharedHolder: holder,
		premiumService:  PremiumService,
	}
}