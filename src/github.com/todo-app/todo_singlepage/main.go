package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database
var sess *mgo.Session

const (
	Hostname       string = "localhost:27017"
	DBname         string = "todolist"
	collectionName string = "todo"
	port           string = ":8080"
)

type (
	todoModel struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		Title     string        `bson:"title"`
		Completed bool          `bson:"completed"`
	}

	todo struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`

	}
)

func init() {
	sess, err := mgo.Dial(Hostname)
	if err != nil {
		panic(err)
	}
	sess.SetMode(mgo.Monotonic, true)
	db = sess.DB(DBname)
}

func homeHandler(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home.tpl",
		gin.H{
			"Title": "Home Page",
		},
	)
}

var router *gin.Engine

func main() {
	defer sess.Close()
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	v1 := router.Group("/v1")
	{

		v1.POST("/", createTodo)
		v1.GET("/", fetchTodos)
		v1.PUT("/{id}", updateTodo)
		v1.DELETE("/{id}", deleteTodo)
	}
	router.Run()
}

func fetchTodos(c *gin.Context) {
	todos := []todoModel{}

	if err := db.C(collectionName).
		Find(bson.M{}).
		All(&todos); err != nil {
		c.JSON(http.StatusProcessing, gin.H{
			"message": "Failed to fetch todo",
			"error":   err,
		})
		return
	}

	todoList := []todo{}
	for _, t := range todos {
		todoList = append(todoList, todo{
			ID:        t.ID.Hex(),
			Title:     t.Title,
			Completed: t.Completed,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": todoList,
	})
}
func createTodo(c *gin.Context) {
	var t todo
	c.BindJSON(&t)
	if t.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The title field is requried",
		})
		return
	} else {
		c.JSON(http.StatusOk, t)
	}

	tm := todoModel{
		ID:        bson.NewObjectId(),
		Title:     t.Title,
		Completed: false,
	}
	if err := db.C(collectionName).Insert(&tm); err != nil {
		c.JSON(http.StatusProcessing, gin.H{
			"message": "Failed to save todo",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo created successfully",
		"todo_id": tm.ID.Hex(),
	})
}
func updateTodo(c *gin.Context) {
	id := bson.ObjectIdHex(c.Param("id"))
	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The id is invalid",
		})
		return
	}

	var t todo
  
	if t.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The title field is requried",
		})
		return
	} else {
		c.JSON(http.StatusOk, t)
	}

	err := db.C(collectionName).UpdateId(bson.M{"_id": bson.ObjectIdHex(id), "title": t.Title, "completed": t.Completed})

	if err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to update todo",
			"error":   err,
		})
		return
	} else {

		c.JSON(http.StatusOK, gin.H{
			"message": "Todo updated successfully",
		})
	}
}
func deleteTodo(c *gin.Context) {
	id := bson.ObjectIdHex(c.Param("id"))
	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The id is invalid",
		})
		return
	}

	if err := db.C(collectionName).RemoveId(bson.ObjectIdHex(id)); err != nil {
		c.JSON(http.StatusProcessing, gin.H{
			"message": "Failed to delete todo",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo deleted successfully",
	})
}
