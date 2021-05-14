package persistence

import (
	"context"
	"dictio-scrapper/model"
	"fmt"

	"github.com/sirupsen/logrus"
)

const STATUS_COLLECTION = "status"

type DB interface {
	Save(entry model.Entry) error
	SaveStatus(status model.Status) error
	LoadStatus() (model.Status, error)
}

type Client interface {
	Database(name string) Database
}

type Database interface {
	Collection(name string) Collection
}

type Collection interface {
	InsertOne(ctx context.Context, data interface{}) (interface{}, error)
	Upsert(ctx context.Context, data interface{}) interface{}
	FindOne(ctx context.Context, data interface{}) interface{}
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

func (db DBImpl) SaveStatus(status model.Status) error {
	err := db.client.
		Database(db.database).
		Collection(STATUS_COLLECTION).
		Upsert(context.TODO(), status)

	if err != nil {
		return fmt.Errorf("Error saving status: %v", err)
	}

	return nil
}

func (db DBImpl) LoadStatus() (model.Status, error) {
	status := model.Status{}

	err := db.client.
		Database(db.database).
		Collection(STATUS_COLLECTION).
		FindOne(context.TODO(), &status)

	if err != nil {
		return model.Status{}, fmt.Errorf("Error loading status: %v", err)
	}

	return status, nil
}
