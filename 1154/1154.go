package main

import (
	"fmt"
)

func main() {

	var x int
	c, soma := 0, 0
	for {
		fmt.Scanln(&x)

		if x < 0 {
			break
		}

		soma += x
		c++
	}

	media := float64(soma) / float64(c)

	fmt.Printf("%.2f\n", media)

}
