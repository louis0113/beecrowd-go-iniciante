package main

import (
	"fmt"
	"math"
)

const Pi = 3.14159

func calcularArea(raio float64) float64 {
	area := Pi * math.Pow(raio, 2)
	return area
}

func main() {
	var r float64
	fmt.Scanln(&r)
	resultado := calcularArea(r)
	fmt.Printf("A=%.4f\n", resultado)
}
