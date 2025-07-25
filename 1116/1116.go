package main

import (
	"fmt"
)

func main() {

	var (
		n, x, y int
		r, z    float64
		s       string
	)

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {

		fmt.Scanf("%d %d ", &x, &y)

		if y == 0 {
			s = "divisao impossivel"
			fmt.Printf("%s\n", s)
		} else if x == 0 {
			z = 0.0
			fmt.Printf("%.1f\n", z)
		} else {
			r = float64(x) / float64(y)
			fmt.Printf("%.1f\n", r)
		}

	}

}
