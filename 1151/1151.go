package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scanln(&n)
	var soma int
	fib := []int{0, 1}

	for c := 1; c < n; c++ {
		soma = fib[c-1] + fib[c]
		fib = append(fib, soma)
	}

	for a := 0; a < n; a++ {
		if a == (n - 1) {
			fmt.Printf("%d\n", fib[a])
		} else {
			fmt.Printf("%d ", fib[a])
		}

	}
}
