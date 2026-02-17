package main

import (
	"fmt"
)

func calcularPorcentagem(sal float64) (pr float64) {
	switch {
	case sal >= 0.00 && sal <= 400.00:
		pr = 15
	case sal > 400.00 && sal <= 800.00:
		pr = 12
	case sal > 800.00 && sal <= 1200.00:
		pr = 10
	case sal > 1200.00 && sal <= 2000.00:
		pr = 7
	default:
		pr = 4
	}
	return
}

func calcularBonus(salario float64) float64 {
	pr := calcularPorcentagem(salario)
	bonus := salario * (pr / 100)
	return bonus
}

func calcularSalario(salario float64) float64 {
	bonus := calcularBonus(salario)
	novo_salario := salario + bonus
	return novo_salario
}

func main() {
	var salario, pr, bonus, novo_salario float64
	fmt.Scanln(&salario)
	pr = calcularPorcentagem(salario)
	novo_salario = calcularSalario(salario)
	bonus = calcularBonus(salario)
	fmt.Printf("Novo salario: %.2f\nReajuste ganho: %.2f\nEm percentual: %d %%\n", novo_salario, bonus, int(pr))
}
