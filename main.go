package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))

	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", router)

	fileServer := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := templates.ExecuteTemplate(w, `index.html`, nil)

	if err != nil {
		panic(err)
	}
}
