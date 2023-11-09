package httpserver

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/configs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type HttpServer struct {
	Router *echo.Echo
	port   string
}

type EnableCors struct {
	IsCorsEnabled bool
}

func NewHttpServer(enableCors ...EnableCors) (*HttpServer, error) {
	port := os.Getenv(configs.PORT_NO)
	e := echo.New()

	// we can use some logger as middleware
	e.Use(middleware.Logger())

	// we can add cors based on some input value
	if len(enableCors) > 1 {
		return nil, errors.New("check parameter passed")
	}

	var enableCorsOption EnableCors

	if len(enableCors) == 1 {
		enableCorsOption = enableCors[0]
	}

	if enableCorsOption.IsCorsEnabled {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPatch, http.MethodHead, http.MethodPut},
			AllowHeaders:     []string{"Content-Type", "X-CSRF-Token", "Authorization"},
			AllowCredentials: true,
			ExposeHeaders:    []string{"Content-Type", "X-CSRF-Token"},
		}))
	}

	return &HttpServer{
		port:   port,
		Router: e,
	}, nil
}

func (s *HttpServer) StartServer() error {
	log.Printf("%s %s", "http server started at", s.port)
	return s.Router.Start(":" + s.port)
}
