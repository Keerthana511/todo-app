package Controllers

import (
	"net/http"

	"./models"
	"github.com/gin-gonic/gin"
)

func fetchAllTodo(context *gin.Context) {
	var todo []models.todo
	err := models.fetchAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
func fetchaTodo(context *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.todo
	err := models.fetchaTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
func createTodo(context *gin.Context) {
	var todo models.todo
	c.BindJSON(&todo)
	err := models.createTodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
func updateTodo(context *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.todo
	err := models.fetchaTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
func deleteTodo(context *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := models.deleteTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
