package main

import (
	"fmt"
)

func main() {

	var x, y int

	fmt.Scanln(&x)
	fmt.Scanln(&y)
	soma := 0

	if x > y {
		for c := x; c >= y; c-- {

			if c%13 != 0 {
				soma += c
			}
		}

	} else {
		for c := x; c <= y; c++ {

			if c%13 != 0 {

				soma += c
			}
		}

	}

	fmt.Println(soma)

}
