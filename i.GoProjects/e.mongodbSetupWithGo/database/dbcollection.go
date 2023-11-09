package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseCollection interface {
	FindOne(ctx context.Context, filter interface{}, document interface{}) error
	FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}) error
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
	Find(ctx context.Context, filter interface{}, options *options.FindOptions, response interface{}) error
	Aggregate(ctx context.Context, pipeline interface{}, response interface{}) error
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Distinct(ctx context.Context, field string, response interface{}) ([]interface{}, error)
	Drop(ctx context.Context) error
	InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
}

type dbcollection struct {
	collection *mongo.Collection
}

func newDatabaseCollection(dbclient *mongo.Database, collection string) DatabaseCollection {
	return &dbcollection{
		collection: dbclient.Collection(collection),
	}
}

func (d *dbcollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return d.collection.DeleteOne(ctx, filter, opts...)
}

func (d *dbcollection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return d.collection.DeleteMany(ctx, filter, opts...)
}

func (d *dbcollection) FindOne(ctx context.Context, filter interface{}, document interface{}) error {
	return d.collection.FindOne(ctx, filter).Decode(document)
}

func (d *dbcollection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}) error {
	result := d.collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (d *dbcollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return d.collection.InsertOne(ctx, document, opts...)
}

func (d *dbcollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return d.collection.UpdateOne(ctx, filter, update, opts...)
}

func (d *dbcollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return d.collection.UpdateMany(ctx, filter, update, opts...)
}

func (d *dbcollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return d.collection.CountDocuments(ctx, filter, opts...)
}

func (d *dbcollection) Find(
	ctx context.Context, filter interface{}, options *options.FindOptions, response interface{}) error {
	results, dberror := d.collection.Find(ctx, filter, options)
	if dberror != nil {
		return dberror
	}
	defer results.Close(ctx)

	//parse the results to array
	return results.All(ctx, response)
}

func (d *dbcollection) Drop(ctx context.Context) error {
	return d.collection.Drop(ctx)
}

func (d *dbcollection) Aggregate(ctx context.Context, pipeline interface{}, response interface{}) error {
	results, dberror := d.collection.Aggregate(ctx, pipeline)
	if dberror != nil {
		return dberror
	}
	defer results.Close(ctx)
	return results.All(ctx, response)
}

func (d *dbcollection) Distinct(
	ctx context.Context, field string, response interface{}) ([]interface{}, error) {
	return d.collection.Distinct(ctx, field, response)
}

func (d *dbcollection) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return d.collection.InsertMany(ctx, documents, opts...)
}
