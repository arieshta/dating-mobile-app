package premium

import (
	"net/http"

	"github.com/arieshta/dating-mobile-app/model/dto"
	"github.com/labstack/echo/v4"
)

func (c *Controller) Purchase(ec echo.Context) error {
	var (
		payload = dto.PurchaseBody{}
	)

	if err := ec.Bind(&payload); err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	premium, err := c.premiumService.Purchase(ec.Get("userid").(string), &payload)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err.Error())
	}

	return ec.JSON(http.StatusOK, premium)
}