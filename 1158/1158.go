package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var x, y, soma int
		fmt.Scanf("%d %d", &x, &y)

		if x%2 == 0 {
			x++
		}

		for j := 0; j < y; j++ {
			soma += x
			x += 2
		}

		fmt.Println(soma)
	}
}
