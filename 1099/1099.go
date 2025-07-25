package main

import (
	"fmt"
)

func main() {

	var x, a, b int

	fmt.Scanln(&x)
	soma := 0

	for c := 0; c < x; c++ {

		fmt.Scanf("%d %d\n", &a, &b)

		start := a
		end := b

		if a > b {
			start = b
			end = a
		}

		for i := start + 1; i < end; i++ {

			if i%2 == 1 {
				soma += i
			}
		}
		fmt.Println(soma)
		soma = 0
	}

}
