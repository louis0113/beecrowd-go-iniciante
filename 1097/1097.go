package main

import (
	"fmt"
)

func main() {

	i, j := 1, 7

	for x := i; x <= 9; x += 2 {

		for a := j; a >= (j - 2); a-- {

			fmt.Printf("I=%d J=%d\n", x, a)
		}

		j += 2
	}
}
