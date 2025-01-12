package handler

import (
	"github.com/arieshta/dating-mobile-app/handler/controllers/matches"
	"github.com/arieshta/dating-mobile-app/handler/controllers/users"
	"github.com/arieshta/dating-mobile-app/model"
)

func RegisterRoutes(h *model.SharedHolder) {
	users.RegisterUserRoutes(h)
	matches.RegisterMatchesRoutes(h)
}