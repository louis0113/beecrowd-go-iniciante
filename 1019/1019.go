package main

import (
	"fmt"
)

func calcularTempo(n int) (h, m, s int) {
	h = n / 3600
	n %= 3600
	m = n / 60
	n %= 60
	s = n
	return
}

func main() {
	var n, horas, minutos, segundos int
	fmt.Scanln(&n)
	horas, minutos, segundos = calcularTempo(n)
	fmt.Printf("%d:%d:%d\n", horas, minutos, segundos)
}
