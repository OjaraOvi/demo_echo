package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"myAppEcho/configs"
	"myAppEcho/routes"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	configs.ConnectDB()
	routes.UserRoute(e)
	e.Logger.Fatal(e.Start(":1323"))
}
