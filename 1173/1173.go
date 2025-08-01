package main

import (
	"fmt"
)

func main() {

	var n [10]int

	for x := 0; x < len(n); x++ {
		if x == 0 {
			fmt.Scanln(&n[0])
		} else {
			n[x] = n[x-1] * 2
		}

		fmt.Printf("N[%d] = %d\n", x, n[x])

	}

}
