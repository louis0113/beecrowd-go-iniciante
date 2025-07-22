package main

import (
	"fmt"
	"math"
)

const Pi = 3.14159

func main() {

	var raio float64

	fmt.Scanln(&raio)

	area := (4.0 / 3) * Pi * math.Pow(raio, 3)

	fmt.Printf("VOLUME = %.3f\n", area)
}
