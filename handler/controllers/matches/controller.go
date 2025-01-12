package matches

import (
	"github.com/arieshta/dating-mobile-app/internal/services/matches"
	"github.com/arieshta/dating-mobile-app/model"
)

type (
	Controller struct {
		sharedHolder *model.SharedHolder

		matchService matches.Service
	}
)

func NewController(holder *model.SharedHolder) *Controller {
	MatchService, err := matches.NewService(holder)
	if err != nil {
		panic(err)
	}

	return &Controller{
		sharedHolder: holder,
		matchService:  MatchService,
	}
}