package main

import (
	"fmt"
)

func calcularDiferenca(a, b, c, d int) int {
	diff := a*b - c*d
	return diff
}

func main() {
	var a, b, c, d int
	fmt.Scanf("%d\n%d\n%d\n%d\n", &a, &b, &c, &d)
	result := calcularDiferenca(a, b, c, d)
	fmt.Println("DIFERENCA =", result)
}
