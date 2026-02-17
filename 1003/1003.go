package main

import (
	"fmt"
)

func sum(x, y int) int {
	soma := x + y
	return soma
}

func main() {
	var a, b int
	fmt.Scanf("%d\n %d\n", &a, &b)
	r := sum(a, b)
	fmt.Println("SOMA =", r)
}
