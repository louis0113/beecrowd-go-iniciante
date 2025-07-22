package main

import (
	"fmt"
)

const mNotas = "nota(s) de R$"
const mMoedas = "moeda(s) de R$"

func main() {

	var dinheiro float64
	var totalNotas, totalMoedas, dinheiroInt int

	fmt.Scanln(&dinheiro)

	notas := [...]int{10000, 5000, 2000, 1000, 500, 200}

	moedas := [...]int{100, 50, 25, 10, 5, 1}

	dinheiroInt = int(dinheiro * 100)

	fmt.Println("NOTAS:")
	for x := 0; x < len(notas); x++ {
		totalNotas = dinheiroInt / notas[x]
		dinheiroInt %= notas[x]

		fmt.Printf("%d %s %.2f\n", totalNotas, mNotas, float64(notas[x])/100)

	}

	fmt.Println("MOEDAS:")
	for y := 0; y < len(moedas); y++ {

		totalMoedas = dinheiroInt / moedas[y]
		dinheiroInt %= moedas[y]

		fmt.Printf("%d %s %.2f\n", totalMoedas, mMoedas, float64(moedas[y])/100)

	}

}
