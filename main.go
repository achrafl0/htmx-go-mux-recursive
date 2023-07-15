package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/clicked", displayState)
	r.HandleFunc("/add/{ID}", addHandler)
	r.HandleFunc("/delete/{ID}", deleteHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
