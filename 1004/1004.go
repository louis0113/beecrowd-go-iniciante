package main

import (
	"fmt"
)

func mult(a, b int) int {
	prod := a * b
	return prod
}

func input(a, b int) (int, int) {
	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%d\n", &b)
	return a, b
}

func main() {
	var x, y int
	x, y = input(x, y)
	resultado := mult(x, y)
	fmt.Println("PROD =", resultado)
}
