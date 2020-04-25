package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("./testdata/raw.html")

	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.html", map[string]interface{}{
			"now": time.Date(2020, 04, 25, 0, 0, 0, 0, time.UTC),
		})
	})

	router.Run(":8080")
}
