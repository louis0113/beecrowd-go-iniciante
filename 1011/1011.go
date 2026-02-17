package main

import (
	"fmt"
	"math"
)

const Pi = 3.14159

func calcularArea(raio float64) float64{
	area := (4.0 / 3) * Pi * math.Pow(raio, 3)
	return area
}

func main() {
	var raio float64
	fmt.Scanln(&raio)
	resultado := calcularArea(raio)
	fmt.Printf("VOLUME = %.3f\n", resultado)
}
