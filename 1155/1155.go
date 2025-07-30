package main

import (
	"fmt"
)

func main() {

	var s, n float64

	for x := 1; x <= 100; x++ {
		n = 1.0 / float64(x)
		s += n

	}

	fmt.Printf("%.2f\n", s)

}
