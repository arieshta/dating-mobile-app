package users

import "github.com/arieshta/dating-mobile-app/model"

// type Handler struct {
// 	Echo *echo.Echo
// }

// func (h *Handler) registerUserRoutes() {
func RegisterUserRoutes(h *model.SharedHolder) {
	controller := NewController(h)

	usersGroup := h.Echo.Group("/users")
	usersGroup.POST("/login", controller.Login)
	usersGroup.POST("/signup", controller.SignUp)
}