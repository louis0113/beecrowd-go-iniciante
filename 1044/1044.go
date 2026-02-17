package main

import (
	"fmt"
)

func verificarMultiplos(x, y int) {
	if y%x == 0 || x%y == 0 {
		fmt.Println("Sao Multiplos")
	} else {
		fmt.Println("Nao sao Multiplos")
	}

}
func main() {
	var a, b int
	fmt.Scanf("%v %v\n", &a, &b)
	verificarMultiplos(a, b)
}
