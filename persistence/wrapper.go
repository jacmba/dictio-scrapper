package persistence

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ClientWrapper struct {
	client *mongo.Client
}

type DbWrapper struct {
	db *mongo.Database
}

type CollectionWrapper struct {
	col *mongo.Collection
}

func NewWrapper(client *mongo.Client) Client {
	return ClientWrapper{client}
}

func (c ClientWrapper) Database(name string) Database {
	db := c.client.Database(name)
	return DbWrapper{db}
}

func (d DbWrapper) Collection(name string) Collection {
	col := d.db.Collection(name)
	return CollectionWrapper{col}
}

func (c CollectionWrapper) InsertOne(ctx context.Context, data interface{}) (interface{}, error) {
	return c.col.InsertOne(ctx, data)
}
