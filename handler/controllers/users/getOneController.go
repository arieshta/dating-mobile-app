package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Controller) GetProfileByUserId(ec echo.Context) error {
	var (
		userId = ec.Param("userId")
	)

	user, err := c.userService.GetOneByUserIdService(userId)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	profile := user.ToProfile()
	if c.premiumService.HasPremium(userId) {
		profile.Verified = true
	}
	return ec.JSON(http.StatusOK, profile)
}