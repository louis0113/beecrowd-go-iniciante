package main

import (
	"fmt"
	"math"
)

func main() {

	var n int

	fmt.Scanln(&n)

	var values [3]float64

	for x := 1; x <= n; x++ {

		for i := 0; i < 3; i++ {
			for c := 0; c < 2; c++ {
				values[i] = math.Pow(float64(x), float64((i + 1)))
				if i == 3 {
					fmt.Printf("%d\n", int(values[i]))

				} else {
					fmt.Printf("%d ", int(values[i]))
				}

			}
		}

	}

}
