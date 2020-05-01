package routes

import (
	"github.com/gin-gonic/gin"
)

func todoroutes() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", Controllers.fetchAllTodo)
		v1.POST("todo", Controllers.createTodo)
		v1.GET("todo/:id", Controllers.fetchaTodo)
		v1.PUT("todo/:id", Controllers.updateTodo)
		v1.DELETE("todo/:id", Controllers.deleteTodo)
	}
}
