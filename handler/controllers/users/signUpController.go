package users

import (
	"net/http"

	"github.com/arieshta/dating-mobile-app/model/dto"
	"github.com/labstack/echo/v4"
)

func (c *Controller) SignUp(ec echo.Context) error {
	var (
		payload = dto.SignUpBody{}
	)
	if err := ec.Bind(&payload); err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}
	user, err := c.userService.SignUpService(&payload)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}
	return ec.JSON(http.StatusOK, user)
}