package main

func todoroutes() {
	router.GET("/", readTodo)
	router.POST("/", createTodo)
	router.PUT("/{id}", updateTodo)
	router.DELETE("/{id}", deleteTodo)
}
