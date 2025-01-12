package matches

import (
	"github.com/arieshta/dating-mobile-app/handler/auth"
	"github.com/arieshta/dating-mobile-app/model"
)

func RegisterMatchesRoutes(h *model.SharedHolder) {
	controller := NewController(h)

	matchesGroup := h.Echo.Group("/matches")
	matchesGroup.Use(auth.Auth)
	matchesGroup.GET("/likes", controller.GetLikes)
	matchesGroup.POST("/likes", controller.LikeOne)
	matchesGroup.GET("/new", controller.GetNewOne)
}