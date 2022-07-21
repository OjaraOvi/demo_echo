package routes

import (
	"github.com/labstack/echo/v4"
	"myAppEcho/controllers"
)

func UserRoute(e *echo.Echo) {
	e.GET("/", controllers.Index)
	e.POST("/user", controllers.CreateUser)
	e.GET("/users", controllers.GetAllUser)
}
