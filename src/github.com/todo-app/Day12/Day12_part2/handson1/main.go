//listen and serve on port 8080 of localhost
package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/stay/", stay)
	http.HandleFunc("/stay.jpg", chien)
	http.ListenAndServe(":8080", nil)
}
func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}
func stay(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("stay.html")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "stay.html", nil)
}
func chien(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "stay.jpg")
}
