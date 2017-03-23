package common

import (
	mgo "gopkg.in/mgo.v2"
)

type mongo struct {
	Tasks *mgo.Collection
}

var DB *mongo

// stsart a MongoDB session
func connectDB() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	DB = &mongo{session.DB("taskdb").C("tasks")}
}
