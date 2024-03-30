package main

import (
	"html/template"
	"net/http"
)

type ViewData struct {
	Name string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Render the index.html template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Get the name parameter from the URL query
	name := r.FormValue("name")

	// Render the hello.html template with the provided name
	tmpl, err := template.ParseFiles("templates/hello.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, ViewData{Name: name})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
