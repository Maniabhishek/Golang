package server

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	port   string
	Router *echo.Echo
}

type HttpServerOptions struct {
	IsCorsEnabled bool
}

func NewHttpServer(context context.Context, options ...HttpServerOptions) (*server, error) {
	var httpServerOptions HttpServerOptions
	if len(options) > 1 {
		return nil, errors.New("length should be 1")
	}

	if len(options) == 1 {
		httpServerOptions = options[0]
	}

	port := os.Getenv("HTTP_PORT")

	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogMethod:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if !strings.HasSuffix(v.URI, "/healthcheck") {
				fmt.Printf("Request URL: %s %s, Status: %d, Latency: %s", v.Method, v.URI, v.Status, v.Latency)
			}
			return nil
		},
	}))

	if httpServerOptions.IsCorsEnabled {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"HEAD", "GET", "POST", "PATCH", "DELETE", "PUT", "OPTIONS"},
			AllowHeaders:     []string{"Content-Type", "X-CSRF-Token", "AccountId", "Authorization", "accountId"},
			AllowCredentials: true,
			ExposeHeaders:    []string{"Content-Type", "X-CSRF-Token"},
		}))
	}

	return &server{
		port:   port,
		Router: e,
	}, nil
}

func (s *server) StartServer() error {
	return s.Router.Start("0.0.0.0:" + s.port)
}
