package main

import (
	"fmt"
)

const pesoA, pesoB, pesoC = 2.0, 3.0, 5.0

func input(x, y, z float64) (a float64, b float64, c float64) {
	fmt.Scanf("%g\n%g\n%g\n", &x, &y, &z)
	a, b, c = x, y, z
	return
}

func mediaPonderada(x, y, z float64) float64 {
	totalPesos := pesoA + pesoB + pesoC
	media := ((x * pesoA) + (y * pesoB) + (z * pesoC)) / totalPesos
	return media
}
func main() {
	var a, b, c float64
	a, b, c = input(a, b, c)
	media := mediaPonderada(a, b, c)
	fmt.Printf("MEDIA = %.1f\n", media)
}
