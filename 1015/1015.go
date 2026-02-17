package main

import (
	"fmt"
	"math"
)

func calcularDistancia(a1, a2, b1, b2 float64) float64{
	dis := math.Sqrt(math.Pow((a2-a1), 2) + math.Pow((b2-b1), 2))
	return dis
}

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scanf("%f %f\n %f %f\n", &x1, &y1, &x2, &y2)
	dis := calcularDistancia(x1,x2,y1,y2)
	fmt.Printf("%.4f\n", dis)
}
