package controllers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	u, err := app.Login(username, password)
	if err != nil {
		return echo.ErrUnauthorized
	}
	// Throws unauthorized error
	// if username != "jon" || password != "shhh!" {
	// 	return echo.ErrUnauthorized
	// }

	// Set custom claims
	claims := &jwtCustomClaims{
		u.FullName(),
		u.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
