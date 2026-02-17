package main

import (
	"fmt"
)

const PesoA, PesoB = 3.5, 7.5

func input(x, y float64) (a float64, b float64) {
	fmt.Scanf("%g\n%g\n", &x, &y)
	a, b = x, y
	return
}

func mediaPonderada(x, y float64) float64 {
	totalPesos := PesoA + PesoB
	media := ((x * PesoA) + (y * PesoB)) / totalPesos
	return media
}

func main() {
	var a, b float64
	a, b = input(a, b)
	r := mediaPonderada(a, b)
	fmt.Printf("MEDIA = %.5f\n", r)
}
