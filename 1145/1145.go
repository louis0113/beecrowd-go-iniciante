package main

import (
	"fmt"
)

func main() {

	var x, y int
	c := 1
	fmt.Scanf("%d %d\n", &x, &y)

	for {

		if c%x != 0 {
			fmt.Printf("%d ", c)
		} else {
			fmt.Printf("%d\n", c)
		}

		c++
		if c == y+1 {
			break
		}

	}

}
