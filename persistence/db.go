package persistence

import (
	"context"
	"dictio-scrapper/model"

	"github.com/sirupsen/logrus"
)

type DB interface {
	Save(entry model.Entry) error
}

type Client interface {
	Database(name string) Database
}

type Database interface {
	Collection(name string) Collection
}

type Collection interface {
	InsertOne(ctx context.Context, data interface{}) (interface{}, error)
}

type DBImpl struct {
	client   Client
	database string
}

func New(client Client, database string) DBImpl {
	return DBImpl{client, database}
}

func (db DBImpl) Save(entry model.Entry) error {
	res, err := db.client.
		Database(db.database).
		Collection(entry.Letters[0]).
		InsertOne(context.TODO(), entry)

	logrus.Infof("Inserted record %v on DB: %v", entry, res)

	return err
}
