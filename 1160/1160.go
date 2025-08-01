package main

import (
	"fmt"
)

func main() {

	var t int

	fmt.Scanln(&t)

	for x := 0; x < t; x++ {

		var pa, pb, c int
		var g1, g2, ppa, ppb float64

		fmt.Scanf("%d %d %f %f", &pa, &pb, &g1, &g2)

		for {

			if pa > pb {
				break
			}

			c++

			ppa = float64(pa) * (g1 / 100)
			pa += int(ppa)
			ppb = float64(pb) * (g2 / 100)
			pb += int(ppb)

		}

		if c > 100 {
			fmt.Println("Mais de 1 seculo.")
		} else {
			fmt.Println(c, "anos.")
		}

	}

}
