package db

import (
	"fmt"

	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

var DB *mgo.Database

func SetupDB() {
	if DB != nil {
		return
	}

	HOST := viper.GetString("mongo.host")
	DATABASE := viper.GetString("mongo.db")

	session, _ := mgo.Dial(fmt.Sprintf("mongodb://%s", HOST))
	defer session.Close()
	DB = session.DB(DATABASE)
}
