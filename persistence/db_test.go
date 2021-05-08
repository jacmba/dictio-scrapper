package persistence

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mockClient struct {
	mock.Mock
}

func (m *mockClient) Database(name string) mongo.Database {
	args := m.Called(name)
	return args.Get(0).(mongo.Database)
}

type mockDb struct {
	mock.Mock
}

func (m *mockDb) Collection(name string) mongo.Collection {
	args := m.Called(name)
	return args.Get(0).(mongo.Collection)
}

type mockCollection struct {
	mock.Mock
}

func (m *mockCollection) InsertOne(ctx context.Context, data bson.M) (interface{}, error) {
	args := m.Called(ctx, data)
	return args.Get(0), args.Error(1)
}
