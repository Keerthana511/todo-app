package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rnd *renderer.Render
var todosCollection *mgo.Collection
var db *mgo.Database

const (
	Hostname       string = "localhost:27017"
	DBname         string = "todolist"
	collectionName string = "todo"
	port           string = ":8080"
)

type Todo struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Title     string
	Completed bool
}
type TodoModel struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Title     string
	Completed bool
}

func init() {
	rnd = renderer.New()
	sess, err := mgo.Dial(Hostname)
	if err != nil {
		panic(err)
	}
	sess.SetMode(mgo.Monotonic, true)
	db = sess.DB(DBname)
}

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	v1 := router.Group("/v1")
	{
		router.GET("/", func(c *gin.Context) {
			c.HTML(
				http.StatusOK,
				"home.tpl",
				gin.H{
					"title": "Home Page",
				},
			)
		})
		v1.POST("/", createTodo)
		v1.GET("/{:id}", fetchTodo)
		v1.GET("/", fetchAllTodo)
		v1.PUT("/{id}", updateTodo)
		v1.DELETE("/{id}", deleteTodo)
	}
	router.Run()
}
func fetchAllTodo(context *gin.Context) {}
func fetchTodo(context *gin.Context)    {}
func createTodo(context *gin.Context)   {}
func updateTodo(context *gin.Context)   {}
func deleteTodo(context *gin.Context)   {}
