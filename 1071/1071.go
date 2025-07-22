package main

import (
	"fmt"
)

func main() {

	var x, y int

	fmt.Scanln(&x)
	fmt.Scanln(&y)

	soma := 0

	for c := (y + 1); c < x; c++ {
		if c%2 == 1 || c%2 == -1 {
			soma += c
		}

	}

	fmt.Println(soma)

}
