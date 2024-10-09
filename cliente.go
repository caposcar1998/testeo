package main

import (
	"fmt"
)

type Cliente struct {
	nombre string
	apellido string
	edad int
}

func (c Cliente) seVa() {
	fmt.Printf("se va")
}