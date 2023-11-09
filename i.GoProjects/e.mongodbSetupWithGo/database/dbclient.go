package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type dbclient struct {
	client   *mongo.Client
	database string
}

type DBClient interface {
	Disconnect(ctx context.Context)
	GetDatabaseName() string
	Collection(collection string) DatabaseCollection
}

func NewDatabaseClient(client *mongo.Client, databasename string) DBClient {
	return &dbclient{
		client:   client,
		database: databasename,
	}
}

func (d *dbclient) Disconnect(ctx context.Context) {
	d.client.Disconnect(ctx)
}

func (d *dbclient) GetDatabaseName() string {
	return d.database
}

func (d *dbclient) Collection(collection string) DatabaseCollection {
	return newDatabaseCollection(d.client.Database(d.database), collection)
}
