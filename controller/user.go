package controller

import (
	"github.com/labstack/echo/v4"
	"myAppEcho/dto"
	"net/http"
)

func GetUser(c echo.Context) error {
	u := &dto.User{
		Name:  "Octavio",
		Email: "octa@lala.com",
	}
	return c.JSON(http.StatusOK, u)
}
