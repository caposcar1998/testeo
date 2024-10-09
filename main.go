package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type person struct {
	name   string
	age    int
	libros [10]int
	casos  map[string]int
}

type customError struct {
	arg     int
	message string
}

func main() {
	fmt.Println("hola mundo")
	variables()

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", homea)
	r.HandleFunc("/produc/{id}", paramss)
	// Start the server
	log.Println("Server starting on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func (e *customError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func conditional(c int) (bool, error) {
	if c > 6 {
		return true, nil
	} else {
		return false, &customError{c, "no puede ser menor"}
	}
}

func variables() {
	var a = "string"
	b := 5
	var d bool = true
	const e = 20
	arreglo := [5]int{1, 2, 3, 4, 5}
	var arreglo2 [10]string
	mapa := make(map[string]int)
	mapa["ho"] = 5
	valor := mapa["ho"]
	fmt.Println(a, b, d, e, arreglo, arreglo2, valor)
	fmt.Println()
}

func sumar(a, b int) int {
	return a + b
}
