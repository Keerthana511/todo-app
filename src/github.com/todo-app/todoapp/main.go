package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		ID        string `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
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
	v1 := router.Group("/v2")
	{
		v1.GET("/home", homeHandler)
		v1.POST("/todo", createTodo)
		v1.GET("/todo", fetchTodos)
		v1.PUT("/todo/:id", updateTodo)
		v1.DELETE("/todo/:id", deleteTodo)
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
	id := c.Param("id")
	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The id is invalid",
		})
		return
	}
	var t todo
	c.BindJSON(&t)
	if t.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The title field is requried",
		})
		return
	}
	if err := db.C(collectionName).
		Update(
			bson.M{"_id": bson.ObjectIdHex(id)},
			bson.M{"title": t.Title, "completed": t.Completed},
		); err != nil {
		c.JSON(http.StatusProcessing, gin.H{
			"message": "Failed to update todo",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo updated successfully",
	})
}
func deleteTodo(c *gin.Context) {
	id := c.Param("id")
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
