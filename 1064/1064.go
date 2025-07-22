package main

import (
	"fmt"
)

func main() {

	var values [6]float64
	c, soma := 0, 0.0
	for x := 0; x < len(values); x++ {

		fmt.Scanln(&values[x])

		if values[x] > 0.00 {
			c++
			soma += values[x]
		}
	}

	media := soma / float64(c)

	fmt.Printf("%d valores positivos\n%.1f\n", c, media)
}
