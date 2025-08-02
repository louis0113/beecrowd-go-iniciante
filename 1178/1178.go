package main

import (
	"fmt"
)

func main() {

	var n [100]float64
	var in float64
	for x := 0; x < len(n); x++ {

		if x == 0 {
			fmt.Scanln(&in)
			n[x] = in
		} else {
			n[x] = n[x-1] / 2.0
		}

		fmt.Printf("N[%d] = %.4f\n", x, n[x])

	}

}
