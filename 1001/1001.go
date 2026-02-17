package main

import (
	"fmt"
)

func soma(x, y int) int {
	sum := x + y
	return sum
}

func main() {
	var a, b int
	fmt.Scanf("%d\n %d\n", &a, &b)
	sum := soma(a, b)
	fmt.Println("X =", sum)
}
