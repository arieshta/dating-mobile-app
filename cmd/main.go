package main

import (
	"fmt"

	"github.com/arieshta/dating-mobile-app/handler"
	"github.com/arieshta/dating-mobile-app/model"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// port := config.LoadEnv().Port

	// app := echo.New()
	// app.Use(middleware.Logger())

	// app.Logger.Fatal(app.Start(port))

	shared := model.SharedHolder{}
	shared.Init()

	shared.Echo.Use(middleware.Logger())

	handler.RegisterRoutes(&shared)

	shared.Echo.Logger.Fatal(shared.Echo.Start(fmt.Sprintf(":%s", shared.Config.Port)))
}