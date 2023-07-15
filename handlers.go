package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var root = Node{
	ID:     uuid.New(),
	childs: []*Node{},
}

var state = State{
	root:     &root,
	allNodes: []*Node{&root},
}

func displayState(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, state.display())
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["ID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := state.add(id)
	if res == -1 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, state.display())
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["ID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := state.delete(id)
	if res == -1 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, state.display())
}
