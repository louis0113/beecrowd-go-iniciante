package main

import (
	"fmt"
)

func main() {

	var n, a int
	soma := 0

	fmt.Scanln(&n)

	for x := 0; x < n; x++ {

		fmt.Scanln(&a)

		for i := 1; i < a; i++ {
			if a%i == 0 {
				soma += i
			}
		}

		if soma == a {
			fmt.Printf("%d eh perfeito\n", a)
		} else {
			fmt.Printf("%d nao eh perfeito\n", a)
		}
		soma = 0
	}
}
