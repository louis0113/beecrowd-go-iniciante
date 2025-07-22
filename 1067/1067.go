package main

import (
	"fmt"
)

func main() {

	var x int

	fmt.Scanln(&x)

	for c := 0; c <= x; c++ {

		if c%2 == 1 {
			fmt.Println(c)
		}

	}

}
