package main

import (
	"fmt"
)

const m = "valor(es)"

func main() {

	var values [5]int
	p, i, po, n := 0, 0, 0, 0
	messages := [...]string{"par(es)", "impar(es)", "positivo(s)", "negativo(s)"}
	for x := 0; x < len(values); x++ {
		fmt.Scanln(&values[x])

		if values[x]%2 == 0 {
			p++
		} else {
			i++
		}

		if values[x] > 0 {
			po++
		} else if values[x] < 0 {
			n++
		}

	}

	r := [...]int{p, i, po, n}

	for y := 0; y < len(r); y++ {
		fmt.Printf("%d %s %s\n", r[y], m, messages[y])
	}
}
