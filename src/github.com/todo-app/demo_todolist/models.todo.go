package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Title     string
	Completed bool
}

var todosCollection *mgo.Collection
var session *mgo.Session
	func init() {
		rnd = renderer.New()
		sess, err := mgo.Dial(hostName)
		checkErr(err)
		sess.SetMode(mgo.Monotonic, true)
		db = sess.DB(dbName)
	}
	

