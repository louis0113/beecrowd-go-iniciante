package main

import (
	"fmt"
)

func main() {

	var a [100]float64

	for x := 0; x < len(a); x++ {
		fmt.Scanln(&a[x])
		if a[x] <= 10 {
			fmt.Printf("A[%d] = %.1f\n", x, a[x])
		}
	}

}
