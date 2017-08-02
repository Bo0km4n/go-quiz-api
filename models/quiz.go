package models

import "gopkg.in/mgo.v2/bson"

type Quiz struct {
	ID         bson.ObjectId `bson:"_id"`
	Name       string        `bson:"name"`
	Difficulty uint          `bson:"difficulty"`
	Subject    string        `bson:"subject"`
	Choices    []string      `bson:"choices"`
	Answer     string        `bson:"answer"`
}
