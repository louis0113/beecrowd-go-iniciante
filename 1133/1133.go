package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	start := x
	end := y

	if x > y {
		start = y
		end = x
	}

	for c := start + 1; c < end; c++ {
		if c%5 == 2 || c%5 == 3 {
			fmt.Println(c)
		}
	}
}
