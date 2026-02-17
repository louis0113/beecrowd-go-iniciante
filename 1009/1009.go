package main

import (
	"fmt"
)

const comissao = 0.15

func calcularSalario(v,s float64) float64{
	bonus := v * comissao
	salario := s + bonus
	return salario
}

func main() {
	var name string
	var salario, vendas float64
	fmt.Scanf("%s\n%g\n%g\n", &name, &salario, &vendas)
	resultado := calcularSalario(vendas, salario)
	fmt.Printf("TOTAL = R$ %.2f\n", resultado)
}
