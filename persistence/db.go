package persistence

import (
	"dictio-scrapper/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type DB interface {
	Save(entry model.Entry) error
}

type DBImpl struct {
	client mongo.Client
}

func New(client mongo.Client) DBImpl {
	return DBImpl{client}
}

func (db DBImpl) Save(entry model.Entry) error {
	return nil
}
