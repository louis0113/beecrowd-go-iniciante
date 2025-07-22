package main

import (
	"fmt"
)

func main() {

	var values [5]int
	c := 0

	for x := 0; x < len(values); x++ {

		fmt.Scanln(&values[x])

		if values[x]%2 == 0 {
			c++
		}

	}

	fmt.Printf("%d valores pares\n", c)
}
