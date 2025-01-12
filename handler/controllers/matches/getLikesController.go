package matches

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *Controller) GetLikes(ec echo.Context) error {
	fmt.Println(ec)
	likes, err := c.matchService.GetLikesService(ec.Get("userid").(string))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ec.JSON(http.StatusNotFound, "No likes found")
		}
		return ec.JSON(http.StatusBadRequest, err.Error())
	}
	return ec.JSON(http.StatusOK, likes)
}