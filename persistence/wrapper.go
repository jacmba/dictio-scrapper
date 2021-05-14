package persistence

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

func (c CollectionWrapper) Upsert(ctx context.Context, data interface{}) interface{} {
	upsert := bson.D{{"$set", data}}
	err := c.col.FindOneAndUpdate(ctx, bson.D{}, upsert).Decode(&bson.M{})

	if err != nil {
		return err
	}

	return nil
}

func (c CollectionWrapper) FindOne(ctx context.Context, data interface{}) interface{} {
	return c.col.FindOne(ctx, bson.D{}).Decode(data)
}
