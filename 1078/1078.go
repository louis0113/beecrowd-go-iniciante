package main

import (
	"fmt"
)

func main() {

	var n, mult int

	fmt.Scanln(&n)

	for x := 1; x < 11; x++ {
		mult = x * n

		fmt.Printf("%d x %d = %d\n", x, n, mult)

	}

}
