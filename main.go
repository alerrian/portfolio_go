package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
)

var templates *template.Template

func main() {
	port := GetPort()
	fmt.Println("Starting Service...")

	fmt.Println("Started...")
	fmt.Println("Listening on Port: " + port)

	templates = template.Must(template.ParseGlob("templates/*.html"))

	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", router)

	fileServer := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	err := http.ListenAndServe(port, router)

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

func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	return ":" + port
}
