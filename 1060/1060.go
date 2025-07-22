package main

import (
	"fmt"
)

func main() {

	var values [6]float64
	c := 0
	for x := 0; x < 6; x++ {
		fmt.Scanln(&values[x])

		if values[x] > 0.0 {
			c++
		}

	}

	fmt.Printf("%d valores positivos\n", c)
}
