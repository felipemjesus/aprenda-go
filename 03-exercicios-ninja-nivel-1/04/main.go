package main

import (
	"fmt"
)

type inteiro int

var x inteiro

func main() {
	fmt.Printf("Valor: %d\nTipo: %T\n", x, x)

	x = 42

	fmt.Println("Valor:", x)
}
