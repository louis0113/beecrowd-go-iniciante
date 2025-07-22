package main

import (
	"fmt"
)

const kml = 12

func main() {

	var horas, kmh int

	fmt.Scanf("%v\n%v\n", &horas, &kmh)

	litros := (float64(horas) * float64(kmh)) / float64(kml)

	fmt.Printf("%.3f\n", litros)
}
