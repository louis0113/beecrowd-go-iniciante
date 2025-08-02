package main

import (
	"fmt"
)

func main() {

	var (
		impar    [5]int
		par      [5]int
		v, in, p int
	)

	for i := 0; i < 15; i++ {

		fmt.Scanln(&v)

		if v%2 == 0 {
			if p == 5 {
				for x := 0; x < p; x++ {
					fmt.Printf("par[%d] = %d\n", x, par[x])
				}
				p = 0
			}
			par[p] = v
			p++
		} else {
			if in == 5 {
				for y := 0; y < in; y++ {
					fmt.Printf("impar[%d] = %d\n", y, impar[y])
				}
				in = 0
			}
			impar[in] = v
			in++
		}
	}

	for z := 0; z < in; z++ {
		fmt.Printf("impar[%d] = %d\n", z, impar[z])
	}

	for h := 0; h < p; h++ {
		fmt.Printf("par[%d] = %d\n", h, par[h])
	}
}
