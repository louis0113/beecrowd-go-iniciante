package main

import (
	"fmt"
)

 func calcularArea(x,y,z float64) float64{
		traArea := ((x + y) * z) / 2
		return traArea
 }

 func calcularPerimetro(x,y,z float64) float64{
		traPeri := x + y + z
		return traPeri
 }

 func validarCondicoes(a,b,c float64){
	 var r float64
	if a < b+c && b < a+c && c < a+b {
		r = calcularPerimetro(a,b,c)
		fmt.Printf("Perimetro = %.1f\n", r)
	} else {
		r = calcularArea(a,b,c)
		fmt.Printf("Area = %.1f\n", r)
	}
 }

func main() {
	var a, b, c float64
	fmt.Scanf("%g %g %g\n", &a, &b, &c)
	validarCondicoes(a,b,c)
}
