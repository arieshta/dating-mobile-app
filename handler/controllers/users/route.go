package users

import (
	"github.com/arieshta/dating-mobile-app/handler/auth"
	"github.com/arieshta/dating-mobile-app/model"
)

func RegisterUserRoutes(h *model.SharedHolder) {
	controller := NewController(h)

	usersGroup := h.Echo.Group("/users", auth.Auth)
	usersGroup.GET("/:userId", controller.GetProfileByUserId)
	usersGroup.POST("/login", controller.Login)
	usersGroup.POST("/signup", controller.SignUp)
	usersGroup.PUT("/:userId", controller.Update)
	usersGroup.DELETE("/:userId", controller.Delete)
}