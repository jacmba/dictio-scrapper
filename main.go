package main

import (
	"context"
	"dictio-scrapper/config"
	"dictio-scrapper/crawler"
	"dictio-scrapper/parser"
	"dictio-scrapper/persistence"
	"strings"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config.LoadConfig()

	getter := crawler.HttpGetterImpl{}
	listParser := parser.NewListParser()
	definitionParser := parser.NewDefinitionParser()

	alphabet := strings.Split(config.GlobalConfig.Alphabet, ",")

	opt := options.Client().ApplyURI(config.GlobalConfig.MongoURL)

	client, err := mongo.Connect(context.TODO(), opt)

	if err != nil {
		logrus.Fatalf("Error connecting to database: %v", err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		logrus.Fatalf("Cannot reach database: %v", err)
	}

	defer client.Disconnect(context.TODO())

	logrus.Info("Successfully connected to database")

	db := persistence.New(client, config.GlobalConfig.Database)

	c := crawler.New(getter, listParser, definitionParser, nil, alphabet)
	c.Process(config.GlobalConfig.URL)
}
