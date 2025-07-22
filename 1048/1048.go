package main

import (
	"fmt"
)

func main() {

	var salario, pr, bonus, novo_salario float64

	fmt.Scanln(&salario)

	switch {

	case salario >= 0.00 && salario <= 400.00:
		pr = 15
	case salario > 400.00 && salario <= 800.00:
		pr = 12
	case salario > 800.00 && salario <= 1200.00:
		pr = 10
	case salario > 1200.00 && salario <= 2000.00:
		pr = 7
	default:
		pr = 4

	}

	bonus = salario * (pr / 100)
	novo_salario = salario + bonus

	fmt.Printf("Novo salario: %.2f\nReajuste ganho: %.2f\nEm percentual: %d %%\n", novo_salario, bonus, int(pr))
}
