package main

import (
	"fmt"
)

func main() {

	var x, soma, c int

	for {

		fmt.Scanln(&x)
		if x == 0 {
			break
		}

		for i := x; i < (x + 10); i++ {

			if i%2 == 0 {
				soma += i
				c++
			}

		}

		fmt.Println(soma)
		soma = 0
	}

}
