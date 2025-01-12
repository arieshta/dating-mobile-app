package users

import (
	"net/http"

	"github.com/arieshta/dating-mobile-app/model/dto"
	"github.com/labstack/echo/v4"
)

func (c *Controller) Login(ec echo.Context) error {
	var (
		payload = dto.LoginBody{}
	)

	if err := ec.Bind(&payload); err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := c.userService.LoginService(&payload)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}
	ec.Set("token", token)
	ec.Response().Writer.Header().Set("Authorization", token)
	return ec.JSON(http.StatusOK, "Success")
	// return ec.String(http.StatusOK, token)
}