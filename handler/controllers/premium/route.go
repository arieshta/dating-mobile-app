package premium

import (
	"github.com/arieshta/dating-mobile-app/handler/auth"
	"github.com/arieshta/dating-mobile-app/model"
)

func RegisterPremiumRoutes(h *model.SharedHolder) {
	controller := NewController(h)

	matchesGroup := h.Echo.Group("/premium", auth.Auth)
	matchesGroup.POST("/purchases", controller.Purchase)
}