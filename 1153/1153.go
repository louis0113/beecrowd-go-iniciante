package main

import (
	"fmt"
)

func main() {

	var x int
	fat := 1

	fmt.Scanln(&x)

	for c := 0; c < x; c++ {
		fat *= (x - c)
	}

	fmt.Println(fat)

}
