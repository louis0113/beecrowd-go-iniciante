package main

import (
	"fmt"
)

const (
	m = "valor(es)"
)

func filtrarNumeros(v [5]int)(arr [4]int){
	var (
		p,i,po,n int
	)
	for x := range 5{
		fmt.Scanln(&v[x])
		if v[x]%2 == 0 {
			p++
		} else {
			i++
		}

		if v[x] > 0 {
			po++
		} else if v[x] < 0 {
			n++
		}
	}
	values := [4]int{p,i,po,n}

	for y := range 4{
		arr[y] = values[y]
	}
	return
}
func main() {
	var values [5]int
	messages := [4]string{"par(es)", "impar(es)", "positivo(s)", "negativo(s)"}
	r := filtrarNumeros(values)
	for y := range 4 {
		fmt.Printf("%d %s %s\n", r[y], m, messages[y])
	}
}
