package controllers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"myAppEcho/configs"
	"myAppEcho/models"
	"net/http"
	"time"
)

func Login(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	username := c.FormValue("username")
	password := c.FormValue("password")
	var user models.User
	defer cancel()
	err := userCollection.FindOne(ctx, bson.M{"name": username}).Decode(&user)

	if username != user.Name {
		return echo.ErrUnauthorized
	}
	error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if error != nil {
		return echo.ErrUnauthorized
	}

	claims := &configs.JwtCustomClaims{
		user.Name,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
