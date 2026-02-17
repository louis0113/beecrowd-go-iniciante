package main

import (
	"fmt"
)

func calcularTempo(x int) int {
	t := x * 2
	return t
}

func main() {
	var x int
	fmt.Scanf("%v\n", &x)
	r := calcularTempo(x)
	fmt.Printf("%v minutos\n", r)
}
