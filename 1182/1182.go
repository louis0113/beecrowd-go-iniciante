package main

import (
	"fmt"
)

func main() {

	var (
		M [12][12]float64
		C int
		T string
		r float64
	)

	fmt.Scanln(&C)
	fmt.Scanln(&T)

	for x := 0; x < 12; x++ {

		for y := 0; y < 12; y++ {

			fmt.Scanln(&M[x][y])

			if y == C {
				r += M[x][y]
			}
		}
	}

	if T == "S" {
		fmt.Printf("%.1f\n", r)
	}

	if T == "M" {
		r /= 12.0
		fmt.Printf("%.1f\n", r)
	}

}
