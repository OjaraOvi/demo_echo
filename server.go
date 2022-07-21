package main

import (
	"github.com/labstack/echo/v4"
	"myAppEcho/configs"
	"myAppEcho/routes"
)

func main() {
	e := echo.New()
	//run database
	configs.ConnectDB()

	routes.UserRoute(e)
	
	e.Logger.Fatal(e.Start(":1323"))
}
