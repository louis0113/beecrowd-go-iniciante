package main

import (
	"fmt"
	"math"
)

func main() {

	var n, r int
	var rf float64
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {

		for x := 0; x < 2; x++ {

			for a := 1; a <= 3; a++ {
				rf = math.Pow(float64(i), float64(a))

				r = int(rf)

				if x == 1 && a > 1 {
					r++
				}

				if a != 3 {
					fmt.Printf("%d ", r)
				} else {
					fmt.Printf("%d\n", r)
				}

			}

		}

	}
}
