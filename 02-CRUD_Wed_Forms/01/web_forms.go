/*
	- Some of the pre-req need to be done, check 02 section = crm_application.go
	- Basic HTML page with http go package
*/
package main

// importing fmt, database/sql, net/http, text/template package
import (
	"log"
	"net/http"
	"text/template"
)

// Home method renders the main.html
func Home(writer http.ResponseWriter, reader *http.Request) {
	var template_html *template.Template
	template_html = template.Must(template.ParseFiles("main.html"))
	template_html.Execute(writer, nil)
}

// main method
func main() {
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}

// Welcome to Web Forms
