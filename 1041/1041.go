package main

import (
	"fmt"
)

func calcularPonto(a,b float64){
	switch {
	case a > 0 && b > 0:
		fmt.Println("Q1")
	case a < 0 && b > 0:
		fmt.Println("Q2")
	case a > 0 && b < 0:
		fmt.Println("Q4")
	case a < 0 && b < 0:
		fmt.Println("Q3")
	case (a < 0 || a > 0) && b == 0:
		fmt.Println("Eixo X")
	case (b < 0 || b > 0) && a == 0:
		fmt.Println("Eixo Y")
	case a == 0 && b == 0:
		fmt.Println("Origem")
	}
} 
func main() {
	var x, y float64
	fmt.Scanf("%g %g\n", &x, &y)
	calcularPonto(x,y)
}
