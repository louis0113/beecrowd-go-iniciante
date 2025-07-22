package main

import (
	"fmt"
)

const p1, p2, p3 = 2.0, 3.0, 5.0
const totalPesos = p1 + p2 + p3

func main() {

	var n int

	fmt.Scanln(&n)

	soma := 0.0
	x := 0
	z := 0
	totallen := n * 3

	values := make([]float64, totallen)
	medias := make([]float64, n)
	pesos := [...]float64{p1, p2, p3}

	for x < n {

		for i := 0; i < 3; i++ {
			if i == 2 {
				fmt.Scanln(&values[z])
			} else {
				fmt.Scan(&values[z])
			}

			soma += values[z] * pesos[i]
		}

		medias[x] = soma / totalPesos

		soma = 0
		z += 3
		x++

	}

	for h := 0; h < n; h++ {
		fmt.Printf("%.1f\n", medias[h])
	}
}
