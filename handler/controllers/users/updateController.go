package users

import (
	"net/http"

	"github.com/arieshta/dating-mobile-app/model/dto"
	"github.com/labstack/echo/v4"
)

func (c *Controller) Update(ec echo.Context) error {
	var (
		requester = ec.Get("userid").(string)
		userId    = ec.Param("userId")
		payload = dto.UpdateBody{}
	)

	if err := ec.Bind(&payload); err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	if userId != requester {
		return ec.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	user, err := c.userService.UpdateService(userId, &payload)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	return ec.JSON(http.StatusOK, user)
}