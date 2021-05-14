package persistence

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestWrapperType(t *testing.T) {
	Convey("Scenario: test DB Wrapper type", t, func() {
		defer func() {
			recover()
		}()
		wrapper := NewWrapper(&mongo.Client{})
		db := DbWrapper{&mongo.Database{}}
		col := CollectionWrapper{&mongo.Collection{}}

		dbRes := wrapper.Database("anyDb")
		colRes := db.Collection("anyCol")
		_, err := col.InsertOne(context.TODO(), bson.M{})
		upRes := col.Upsert(context.TODO(), bson.M{})
		findRes := col.FindOne(context.TODO(), bson.M{})

		So(dbRes, ShouldBeNil)
		So(colRes, ShouldBeNil)
		So(err, ShouldNotBeNil)
		So(upRes, ShouldNotBeNil)
		So(findRes, ShouldNotBeNil)
	})
}
