package main

import (
	"fmt"
)

func main() {

	i, j := 1, 7

	for y := i; y <= 9; y += 2 {
		for x := j; x >= 5; x-- {
			fmt.Printf("I=%d J=%d\n", y, x)
		}
	}
}
