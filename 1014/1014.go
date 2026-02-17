package main

import (
	"fmt"
)

func bebidaCarro(a int, b float32) float32{
	kml := float32(a) / b
	return kml
}

func main() {
	var x int
	var y float32
	fmt.Scanf("%d\n%f\n ", &x, &y)
	resultado := bebidaCarro(x,y)
	fmt.Printf("%.3f km/l\n", resultado)
}
