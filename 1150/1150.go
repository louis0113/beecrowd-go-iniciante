package main

import (
	"fmt"
)

func main() {

	var x, z, c, soma int
	fmt.Scanln(&x)

	for {

		if z > x {
			break
		}

		fmt.Scanln(&z)
	}

	for a := x; a <= z; a++ {
		soma += a
		c++
		if soma > z {
			break
		}
	}

	fmt.Println(c)
}
