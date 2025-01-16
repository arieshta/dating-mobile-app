package matches

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Controller) PassOne(ec echo.Context) error {
	matchId := ec.QueryParam("matchid")
	feed, err := c.matchService.PassOneService(ec.Get("userid").(string),matchId)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}
	return ec.JSON(http.StatusOK, feed)
}
