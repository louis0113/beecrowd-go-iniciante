package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Scanln(&n)

	for x := 1; x <= 10000; x++ {

		if x%n == 2 {
			fmt.Println(x)
		}
	}

}
