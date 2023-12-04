package controllers

import (
	"github.com/akriventsev/sample-rest-layered/source/internal/app/application"
	"github.com/golang-jwt/jwt/v5"
)

var secret string
var app application.IApplication

type jwtCustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.RegisteredClaims
}

func InitSecret(s string) {
	secret = s
}
func InitApp(a application.IApplication) {
	app = a
}
