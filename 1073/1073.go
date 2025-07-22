package main

import (
	"fmt"
	"math"
)

const msg = "^2 ="

func main() {

	var n int

	fmt.Scanln(&n)
	pot := 0.0
	for x := 1; x <= n; x++ {

		if x%2 == 0 {
			pot = math.Pow(float64(x), 2)

			fmt.Printf("%d%s %d\n", x, msg, int(pot))
		}

	}

}
