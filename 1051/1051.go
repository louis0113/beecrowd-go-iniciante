package main

import (
	"fmt"
)

func calcularSalario(salario float64) (resultado string) {
	var total float64
	switch {
	case salario >= 0.00 && salario <= 2000.00:
		resultado = fmt.Sprintln("Isento")
	case salario > 2000.00 && salario <= 3000.00:
		total = (salario - 2000.00) * 0.08
	case salario > 3000.00 && salario <= 4500.00:
		total = (1000.00 * 0.08) + ((salario - 3000.00) * 0.18)
	default:
		total = (1000.00 * 0.08) + (1500.00 * 0.18) + ((salario - 4500.00) * 0.28)
	}
	resultado = fmt.Sprintf("R$ %.2f\n", total)
	return
}

func main() {
	var salario float64
	fmt.Scanln(&salario)
	r := calcularSalario(salario)
	fmt.Println(r)
}
