package premium

// import "github.com/labstack/echo/v4"

// func (c *Controller) Purchase(ec echo.Context) error {
// 	likes, err := c.premiumService.Purchase(ec.Get("userid").(string))
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return ec.JSON(http.StatusNotFound, "No likes found")
// 		}
// 		return ec.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	return ec.JSON(http.StatusOK, likes)
// }