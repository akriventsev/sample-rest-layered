package controllers

import (
	"net/http"
	"time"

	"github.com/akriventsev/sample-rest-layered/source/internal/app/entities"
	"github.com/labstack/echo/v4"
)

type signupRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Married   bool      `json:"married"`
	Password  string    `json:"password"`
	Login     string    `json:"login"`
	Birthday  time.Time `json:"birthday"`
}

type signupResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Married   bool      `json:"married"`
	Login     string    `json:"login"`
	Birthday  time.Time `json:"birthday"`
}

func SignUp(c echo.Context) error {
	// username := c.FormValue("username")
	// password := c.FormValue("password")

	sr := signupRequest{}
	if err := c.Bind(&sr); err != nil {
		return echo.ErrBadRequest
	}

	u := entities.User{
		FirstName: sr.FirstName,
		LastName:  sr.LastName,
		Married:   sr.Married,
		Password:  sr.Password,
		Login:     sr.Login,
		Birthday:  sr.Birthday,
	}

	if err := u.Validate(); err != nil {
		return err
	}

	if user, err := app.CreateUser(u); err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, signupResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Married:   user.Married,
			Login:     user.Login,
			Birthday:  user.Birthday,
		})
	}
}
