package main

import (
	"github.com/labstack/echo/v4"
	"myAppEcho/configs"
	"myAppEcho/controller"
)

func main() {
	e := echo.New()
	//run database
	configs.ConnectDB()
	e.GET("/", controller.Index)
	e.GET("/users", controller.GetUser)
	e.Logger.Fatal(e.Start(":1323"))
}
