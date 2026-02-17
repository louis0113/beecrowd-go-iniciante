package main

import (
	"fmt"
)

func calcularSalario(h int, s float64) float64{
	totalSalary := float64(h) * s
	return totalSalary
}

func main() {
	var numFunc, hour int
	var salary float64
	fmt.Scanf("%d\n%d\n%g\n", &numFunc, &hour, &salary)
	r := calcularSalario(hour, salary)
	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", numFunc, r)
}
