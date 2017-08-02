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
	PORT := viper.GetString("mongo.port")
	DATABASE := viper.GetString("mongo.db")

	session, _ := mgo.Dial(fmt.Sprintf("mongodb://%s"))
	defer session.Close()
	DB = session.DB("quiz_ios")
}
