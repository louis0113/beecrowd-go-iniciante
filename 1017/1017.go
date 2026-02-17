package main

import (
	"fmt"
)

const kml = 12

func calcularLitros(h, k int) float64{
	litros := (float64(h) * float64(k)) / float64(kml)
	return litros
}

func main() {
	var horas, kmh int
	fmt.Scanf("%v\n%v\n", &horas, &kmh)
	r := calcularLitros(horas,kmh)
	fmt.Printf("%.3f\n", r)
}
