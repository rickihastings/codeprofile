package modules

import (
	"labix.org/v2/mgo"
)

func DbConnect() *mgo.Database {
	mongo, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	return mongo.DB("codeprofile")
}