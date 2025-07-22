package main

import (
	"fmt"
	"math"
)

const Pi = 3.14159

func main() {

	var x, y, z float64

	fmt.Scanf("%g %g %g\n", &x, &y, &z)

	tri, cir, tra, qua, ret := areas(x, y, z)
	resultados := [...]float64{tri, cir, tra, qua, ret}
	nomes := [...]string{"TRIANGULO", "CIRCULO", "TRAPEZIO", "QUADRADO", "RETANGULO"}

	for i := 0; i < 5; i++ {

		fmt.Printf("%s: %.3f\n", nomes[i], resultados[i])
	}
}

func areas(a, b, c float64) (tri, cir, tra, qua, ret float64) {

	tri = (a * c) / 2
	cir = Pi * math.Pow(c, 2)
	tra = ((a + b) * c) / 2
	qua = math.Pow(b, 2)
	ret = a * b
	return

}
