package main

import (
	"fmt"
)

func main() {

	soma := 0.0
	in := 1.0
	for c := 1; c <= 39; c += 2 {

		soma += float64(c) / float64(in)
		in *= 2
	}

	fmt.Printf("%.2f\n", soma)
}
