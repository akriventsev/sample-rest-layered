package controllers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/akriventsev/sample-rest-layered/source/internal/app/entities"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type orderItem struct {
	ProductID string `json:"product_id"`
	Quantity  uint64 `json:"quantity"`
	Price     uint64 `json:"price"`
}
type orderRequest struct {
	Items []orderItem `json:"items"`
}

type orderResponse struct {
	ID    string      `json:"id"`
	Items []orderItem `json:"items"`
}

func Order(c echo.Context) error {
	rawToken := c.Get("user")
	slog.Info("get token", "token", rawToken, "type", fmt.Sprintf("%T", rawToken))

	token, ok := rawToken.(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		slog.Error("JWT token missing or invalid")
		return errors.New("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`

	if !ok {
		slog.Error("failed to cast claims as jwt.MapClaims")
		return errors.New("failed to cast claims as jwt.MapClaims")
	}

	raw_user_id, ok := claims["id"]

	if !ok {
		slog.Error("failed to get user id from token", "claims", claims)
		return errors.New("failed to get user id from token")
	}

	user_id, ok := raw_user_id.(string)

	if !ok {
		return errors.New("failed to cast user id from token")
	}

	or := orderRequest{}
	if err := c.Bind(&or); err != nil {
		return echo.ErrBadRequest
	}

	o := entities.Order{
		UserID: user_id,
		Items:  []entities.OrderItem{},
	}

	for _, oi := range or.Items {
		o.Items = append(o.Items, entities.OrderItem{
			ProductID: oi.ProductID,
			Quantity:  oi.Quantity,
		})
	}

	if order, err := app.CreateOrder(o); err != nil {
		return err
	} else {
		orderResp := orderResponse{
			ID:    order.ID,
			Items: []orderItem{},
		}

		for _, v := range order.Items {
			orderResp.Items = append(orderResp.Items, orderItem{
				ProductID: v.ProductID,
				Quantity:  v.Quantity,
				Price:     v.Price,
			})
		}

		return c.JSON(http.StatusOK, orderResp)
	}
}
