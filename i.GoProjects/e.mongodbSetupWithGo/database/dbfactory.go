package database

import (
	"context"
	"log"
	"os"

	"github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseFactory interface {
	NewDBConnection(ctx context.Context) (DBClient, error)
}

type databasefactory struct {
}

func NewDatabaseFactory() DatabaseFactory {
	return &databasefactory{}
}

func (d *databasefactory) NewDBConnection(ctx context.Context) (DBClient, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// defer cancel()

	mongo_uri := os.Getenv(configs.MONGO_URI)
	databasename := os.Getenv(configs.DB_NAME)
	coptions := options.Client().ApplyURI(mongo_uri)

	mongo_auth := os.Getenv(configs.MONGO_AUTH)
	if mongo_auth == "true" {
		credentials := options.Credential{
			Username: os.Getenv(configs.MONGO_USERNAME),
			Password: os.Getenv(configs.MONGO_PASSWORD),
		}
		coptions.SetAuth(credentials)
	}
	client, err := mongo.Connect(ctx, coptions)
	if err != nil {
		return nil, err
	}

	log.Printf("%s %s", "connected to the database", databasename)
	return NewDatabaseClient(client, databasename), nil
}
