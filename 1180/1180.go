package main

import (
	"fmt"
)

func main() {
	var n, menor, i int

	fmt.Scanln(&n)

	x := make([]int, n)

	for c := 0; c < len(x); c++ {

		fmt.Scan(&x[c])
	}

	i = 1

	menor = x[0]

	for i < len(x) {

		if menor > x[i] {
			menor = x[i]
		}
		i++
	}

	for a := 0; a < len(x); a++ {

		if x[a] == menor {
			fmt.Println("Menor valor:", x[a])
			fmt.Println("Posicao:", a)
		}
	}

}
