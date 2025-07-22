package main

import (
	"fmt"
)

const PesoA = 3.5
const PesoB = 7.5

func main() {

	var a, b float64

	fmt.Scanf("%g\n%g\n", &a, &b)

	media := ((a * PesoA) + (b * PesoB)) / (PesoA + PesoB)

	fmt.Printf("MEDIA = %.5f\n", media)
}
