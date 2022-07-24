package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"myAppEcho/configs"
	"myAppEcho/controllers"
)

func UserRoute(e *echo.Echo) {
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &configs.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	// Restricted group
	r := e.Group("/api")
	r.Use(middleware.JWTWithConfig(config))
	e.GET("/", controllers.Index)
	e.POST("/login", controllers.Login)
	r.POST("/user", controllers.CreateUser)
	r.GET("/users", controllers.GetAllUser)
	r.GET("/user/:userId", controllers.GetAUser)
	r.PUT("/user/:userId", controllers.EditAUser)
	r.DELETE("/user/:userId", controllers.DeleteAUser)
}
