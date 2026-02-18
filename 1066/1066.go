package main

import (
	"fmt"
)

const m = "valor(es)"

func main() {
	var values [5]int
	p, i, po, n := 0, 0, 0, 0
	messages := [4]string{"par(es)", "impar(es)", "positivo(s)", "negativo(s)"}
	for x := range 5{
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
	r := [4]int{p, i, po, n}
	for y := range 4 {
		fmt.Printf("%d %s %s\n", r[y], m, messages[y])
	}
}
