package persistence

import (
	"context"
	"dictio-scrapper/model"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

type mockClient struct {
	mock.Mock
}

func (m *mockClient) Database(name string) Database {
	args := m.Called(name)
	return args.Get(0).(Database)
}

type mockDb struct {
	mock.Mock
}

func (m *mockDb) Collection(name string) Collection {
	args := m.Called(name)
	return args.Get(0).(Collection)
}

type mockCollection struct {
	mock.Mock
}

func (m *mockCollection) InsertOne(ctx context.Context, data interface{}) (interface{}, error) {
	args := m.Called(ctx, data)
	return args.String(0), args.Error(1)
}

func (m *mockCollection) Upsert(ctx context.Context, data interface{}) interface{} {
	args := m.Called(ctx, data)
	return args.Get(0)
}

func (m *mockCollection) FindOne(ctx context.Context, data interface{}) interface{} {
	args := m.Called(ctx, data)
	return args.Get(0)
}

func TestDBStorage(t *testing.T) {
	Convey("Scenario: Test DB storage", t, func() {
		client := new(mockClient)
		database := new(mockDb)
		collection := new(mockCollection)

		client.On("Database", "myDb").Return(database)
		database.On("Collection", mock.AnythingOfType("string")).Return(collection)
		collection.On("InsertOne", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("model.Entry")).Return("ok", nil)
		collection.On("Upsert", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("model.Status")).Return(nil)
		collection.On("FindOne", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("*model.Status")).Run(func(args mock.Arguments) {
			status := args.Get(1).(*model.Status)
			status.Letter = "M"
			status.Word = "mock"
			status.Timestamp = "now"
		}).Return(nil)

		Convey("Given a DB instance", func() {
			db := New(client, "myDb")

			Convey("When Save method is called", func() {
				data := model.Entry{
					Word:       "foo",
					Definition: "lorem ipsum dolor sit amet",
					Letters:    []string{"f"},
				}
				err := db.Save(data)

				Convey("Then should be no errors and DB methods must be invoked", func() {
					So(err, ShouldBeNil)
					client.AssertCalled(t, "Database", "myDb")
					database.AssertCalled(t, "Collection", "f")
					collection.AssertCalled(t, "InsertOne", mock.AnythingOfType("*context.emptyCtx"), data)
				})
			})

			Convey("When SaveStatus method is called", func() {
				data := model.Status{
					Letter:    "G",
					Word:      "gambitero",
					Timestamp: "ayer",
				}
				err := db.SaveStatus(data)

				Convey("Then should be no errors and DB methods must be invoked", func() {
					So(err, ShouldBeNil)
					client.AssertCalled(t, "Database", "myDb")
					database.AssertCalled(t, "Collection", "status")
					collection.AssertCalled(t, "Upsert", mock.AnythingOfType("*context.emptyCtx"), data)
				})
			})

			Convey("When LoadStatus method is called", func() {
				data, err := db.LoadStatus()

				Convey("Then should be no errors and DB methods must be invoked", func() {
					So(err, ShouldBeNil)
					So(data, ShouldNotBeNil)
					So(data.Letter, ShouldEqual, "M")
					So(data.Word, ShouldEqual, "mock")
					So(data.Timestamp, ShouldEqual, "now")
				})
			})
		})
	})
}
