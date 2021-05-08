package persistence

import (
	"context"
	"dictio-scrapper/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type DB interface {
	Connect() error
	Save(entry model.Entry) error
}

type DBImpl struct {
	client mongo.Client
	ctx    context.Context
}

func New(client mongo.Client) DBImpl {
	return DBImpl{client, context.TODO()}
}

func (db DBImpl) Connect() error {
	return nil
}

func (db DBImpl) Save(entry model.Entry) error {
	return nil
}
