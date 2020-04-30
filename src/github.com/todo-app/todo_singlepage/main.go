package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
func fetchAllTodo(context *gin.Context) {
	var todos []Todo
	err := todosCollection.Find(nil).All(&todos)
	if err != nil {
		log.Fatal(err)
	}

	if len(todos) <= 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "no todo found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   todos,
	})
}
func fetchTodo(context *gin.Context) {
	todo := Todo{}
	id := bson.ObjectIdHex(context.Param("id"))
	err := todosCollection.FindId(id).One(&todo)

	if err != nil || todo == (Todo{}) {
		fmt.Println("Error: " + err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "todo not found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   todo,
	})
}
func createTodo(context *gin.Context) {
	title := context.PostForm("Title")
	completed, _ := strconv.ParseBool(context.PostForm("Completed"))
	var todo = Todo{bson.NewObjectId(), title, completed}
	fmt.Println("" + todo.Title + " completed: " + strconv.FormatBool(todo.Completed))
	err := todosCollection.Insert(&todo)
	if err != nil {
		log.Fatal(err)
	}

	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "todo item created successfully",
	})
}
func updateTodo(context *gin.Context) {
	id := bson.ObjectIdHex(context.Param("id"))
	title := context.PostForm("title")
	completed, _ := strconv.ParseBool(context.PostForm("completed"))

	err := todosCollection.UpdateId(id, bson.M{"title": title, "completed": completed})

	fmt.Printf("completed: %t\n\n", completed)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "todo not found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo updated successfully!",
	})
}
func deleteTodo(context *gin.Context) {
	id := bson.ObjectIdHex(context.Param("id"))

	fmt.Printf("id: %v", id)

	err := todosCollection.RemoveId(id)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "todo not found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo deleted successfully!",
	})
}
