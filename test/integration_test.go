// +build integration

package test

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestApplicationIntegration(t *testing.T) {
	Convey("Scenario: test application behavior", t, func() {
		opt := options.Client().ApplyURI("mongodb://mongo:27017")
		client, err := mongo.Connect(context.TODO(), opt)

		So(err, ShouldBeNil)

		defer client.Disconnect(context.TODO())

		logrus.Info("Connected to MongoDB. Checking ping...")
		err = client.Ping(context.TODO(), nil)
		So(err, ShouldNotBeNil)

		logrus.Info("Ping suceeded. Checking available databases...")
		dbs, err := client.ListDatabaseNames(context.TODO(), bson.D{})

		logrus.Infof("Databases found in test Mongo: %v", dbs)

		So(len(dbs), ShouldBeGreaterThan, 1)

		col := client.Database("dictio").Collection("a")
		count, err := col.CountDocuments(context.TODO(), bson.D{})

		So(count, ShouldBeGreaterThan, 3)
	})
}
