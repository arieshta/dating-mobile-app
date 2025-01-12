package matches

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Controller) GetNewOne(ec echo.Context) error {
	feed, err := c.matchService.GetNewOneFeedService(ec.Get("userid").(string))
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}
	return ec.JSON(http.StatusOK, feed)
}
