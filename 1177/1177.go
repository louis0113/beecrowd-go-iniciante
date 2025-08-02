package main

import (
	"fmt"
)

func main() {

	var t, c int
	var n [1000]int
	fmt.Scanln(&t)
	total := len(n) / t
	r := len(n) % t
	for x := 0; x < int(total); x++ {
		for i := 0; i <= t-1; i++ {
			n[c] = i
			fmt.Printf("N[%d] = %d\n", c, n[c])
			c++
		}
	}

	if r > 0 {
		for a := 0; a < r; a++ {
			n[c] = a
			fmt.Printf("N[%d] = %d\n", c, n[c])
			c++
		}
	}
}
