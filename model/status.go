package model

type Status struct {
	Letter    string `bson:"letter"`
	Word      string `bson:"word"`
	Timestamp string `bson:"timestamp"`
}
