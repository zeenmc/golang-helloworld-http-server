package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "homepage.tmpl")

}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {

		fmt.Println("Error parsing tempalte: ", err)
		return
	}

}

func main() {

	http.HandleFunc("/", Home)

	_ = http.ListenAndServe(":5555", nil)

}
