package main

import (
	"fmt"
)

func main() {

	fib := [61]uint64{0, 1}

	var t, n, x, a uint64

	fmt.Scanln(&t)

	for x = 0; x < t; x++ {

		fmt.Scanln(&n)

		if n >= 2 {
			for a = 2; a <= n; a++ {

				fib[a] = fib[a-1] + fib[a-2]
			}

		}

		fmt.Printf("Fib(%d) = %d\n", n, fib[n])
	}
}
