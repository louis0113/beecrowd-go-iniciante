package main

import (
	"fmt"
)

func main() {

	var n [20]int
	var temp int

	total := len(n)

	for x := 0; x < total; x++ {
		fmt.Scanln(&n[x])
	}

	for i := 0; i < total; i++ {

		if i == 10 {
			break
		}
		temp = n[i]
		n[i] = n[total-i-1]
		n[total-i-1] = temp

	}

	for c := 0; c < total; c++ {
		fmt.Printf("N[%d] = %d\n", c, n[c])
	}

}
