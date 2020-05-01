package main

import (
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
)

var todosCollection *mgo.Collection
var session *mgo.Session

const (
	Hostname       string = "localhost:27017"
	DBname         string = "todolist"
	collectionName string = "todo"
	port           string = ":8080"
)

func init() {
	rnd = renderer.New()
	sess, err := mgo.Dial(hostName)
	checkErr(err)
	sess.SetMode(mgo.Monotonic, true)
	db = sess.DB(dbName)
}
