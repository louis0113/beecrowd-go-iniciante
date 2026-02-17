package main

import (
	"fmt"
)

const mensagem = "nota(s) de R$"

func mostrarNotas(quan int, tipo [7]int) {
	fmt.Println(quan)
	for x := range len(tipo) {
		total := quan / tipo[x]
		quan %= tipo[x]
		fmt.Printf("%v %s %v,00\n", total, mensagem, tipo[x])
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	notas := [7]int{100, 50, 20, 10, 5, 2, 1}
	mostrarNotas(n, notas)
}
