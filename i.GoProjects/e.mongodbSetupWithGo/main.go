package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/apis"
	"github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/database"
	httpserver "github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/http-server"
	"github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/internal/db"
	"github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctxt := context.Background()

	dbfactory := database.NewDatabaseFactory()
	client, err := dbfactory.NewDBConnection(ctxt)
	if err != nil {
		panic(err)
	}

	dbFactory := db.NewDbFactory(client)

	servicefactory := services.NewServiceFactory(dbFactory)

	server, err := httpserver.NewHttpServer(httpserver.EnableCors{IsCorsEnabled: true})
	if err != nil {
		panic(err)
	}

	go func() {
		group := server.Router.Group("/api/v1")

		api := apis.NewHealthCheckAPI()
		postapi := apis.NewPostAPI(servicefactory)
		group.GET("/health", api.HealthAPI)
		group.POST("/post", postapi.SavePost)

		server.StartServer()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Router.Shutdown(ctx); err != nil {
		server.Router.Logger.Fatal(err)
	}
}
