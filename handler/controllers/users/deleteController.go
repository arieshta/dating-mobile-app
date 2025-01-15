package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Controller) Delete(ec echo.Context) error {
	var (
		requester = ec.Get("userid").(string)
		userId    = ec.Param("userId")
	)

	if userId != requester {
		return ec.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	
	err := c.userService.DeleteService(userId)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}
	return ec.JSON(http.StatusOK, "Success")
}