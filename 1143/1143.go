package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	var r float64

	fmt.Scanln(&n)

	for x := 1; x <= n; x++ {

		for i := 1; i <= 3; i++ {

			r = math.Pow(float64(x), float64(i))

			if i == 3 {
				fmt.Printf("%d\n", int(r))
			} else {
				fmt.Printf("%d ", int(r))
			}
		}
	}

}
