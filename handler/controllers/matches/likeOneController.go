package matches

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Controller) LikeOne(ec echo.Context) error {
	matchId := ec.QueryParam("matchid")
	likes, err := c.matchService.LikeOneService(ec.Get("userid").(string),matchId)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}
	return ec.JSON(http.StatusOK, likes)
}
