package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/akriventsev/sample-rest-layered/source/internal/app/application"
	"github.com/akriventsev/sample-rest-layered/source/internal/transport/http/controllers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// Option is a functional option type that allows us to configure the Client.
type Option func(*Transport)

type Config struct {
	Port          uint16
	ListenAddress string
	JwtSecret     string
}

type Transport struct {
	app    application.IApplication
	config Config
}

func WithConfig(cfg Config) Option {
	return func(c *Transport) {
		c.config = cfg
	}
}

// NewClient creates a new HTTP client with default options.
func NewTransport(app application.IApplication, options ...Option) (*Transport, error) {
	if app == nil {
		return nil, fmt.Errorf("cannot start with nil app")
	}

	transport := &Transport{
		app: app,
		config: Config{
			ListenAddress: "0.0.0.0",
			Port:          1323,
			JwtSecret:     "secret",
		},
	}

	// Apply all the functional options to configure the client.
	for _, opt := range options {
		opt(transport)
	}

	controllers.InitSecret(transport.config.JwtSecret)
	controllers.InitApp(transport.app)

	return transport, nil
}

type AppErrorResp struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func ErrorHandler(err error, c echo.Context) {
	c.JSON(http.StatusOK, AppErrorResp{
		Success: false,
		Error:   err.Error(),
	})
}

func (t Transport) Start(ctx context.Context) error {
	e := echo.New()

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	})

	e.POST("/login", controllers.Login)
	e.POST("/signup", controllers.SignUp)

	e.POST("/order", controllers.Order, echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(t.config.JwtSecret),
	}))

	e.HTTPErrorHandler = ErrorHandler

	err := e.Start(fmt.Sprintf("%s:%d", t.config.ListenAddress, t.config.Port))

	if err != nil {
		return err
	}

	return nil
}
