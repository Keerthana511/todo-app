package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/ginS"
)

func main() {
	ginS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.Run()
}
