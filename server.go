package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homea(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola mundo")
}

func paramss(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product := vars["id"]
	fmt.Fprint(w, product)
}
