package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"myAppEcho/api/login"
	"myAppEcho/api/user"
	"myAppEcho/configs"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	configs.ConnectDB()
	route(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func route(e *echo.Echo) {
	userController := &user.Handler{}
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &configs.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	// Restricted group
	r := e.Group("/api")
	r.Use(middleware.JWTWithConfig(config))
	//e.GET("/", Index)
	e.POST("/login", login.Login)
	r.POST("/user", userController.CreateUser)
	r.GET("/users", user.GetAllUser)
	r.GET("/user/:userId", user.GetAUser)
	r.PUT("/user/:userId", user.EditAUser)
	r.DELETE("/user/:userId", user.DeleteAUser)
}
