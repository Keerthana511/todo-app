package main

import (
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
)

var todosCollection *mgo.Collection
var session *mgo.Session

func init() {
	rnd = renderer.New()
	sess, err := mgo.Dial(hostName)
	checkErr(err)
	sess.SetMode(mgo.Monotonic, true)
	db = sess.DB(dbName)
}
