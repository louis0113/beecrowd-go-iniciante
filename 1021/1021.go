package main

import (
	"fmt"
)

const mNotas = "nota(s) de R$"
const mMoedas = "moeda(s) de R$"

func mostrarValores(d float64, n, m [6]int) {
	dinheiroInt := int(d * 100)
	fmt.Println("NOTAS:")
	for x := range len(n) {
		totalNotas := dinheiroInt / n[x]
		dinheiroInt %= n[x]
		fmt.Printf("%d %s %.2f\n", totalNotas, mNotas, float64(n[x])/100)
	}
	fmt.Println("MOEDAS:")
	for y := range len(m) {
		totalMoedas := dinheiroInt / m[y]
		dinheiroInt %= m[y]
		fmt.Printf("%d %s %.2f\n", totalMoedas, mMoedas, float64(m[y])/100)
	}
}

func main() {
	var dinheiro float64
	fmt.Scanln(&dinheiro)
	notas := [6]int{10000, 5000, 2000, 1000, 500, 200}
	moedas := [6]int{100, 50, 25, 10, 5, 1}
	mostrarValores(dinheiro, notas, moedas)
}
