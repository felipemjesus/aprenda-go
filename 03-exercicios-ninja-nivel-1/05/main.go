package main

import (
	"fmt"
)

type inteiro int

var x inteiro
var y int

func main() {
	fmt.Printf("Valor de x: %d\nTipo de x: %T\n", x, x)

	x = 42
	fmt.Println("Valor de x:", x)

	y = int(x)
	fmt.Printf("Valor de y: %d\nTipo de y: %T\n", y, y)
}
