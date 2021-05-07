package model

// Entry data structure of word definition for persistence
type Entry struct {
	Word       string   `bson:"word"`
	Definition string   `bson:"definition"`
	Letters    []string `bson:"letters"`
}
