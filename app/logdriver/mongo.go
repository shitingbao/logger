package logdriver

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func OpenMongo(driver, database string) (*mongo.Database, error) {
	db, err := openMongo(driver, database)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func openMongo(driver, database string) (*mongo.Database, error) {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(driver))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return client.Database(database), nil
}
