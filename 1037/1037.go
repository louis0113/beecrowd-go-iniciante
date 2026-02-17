package main

import (
	"fmt"
)

const mensagem = "Intervalo"

func calcularIntervalo(a float64){
	switch {
	case a >= 0.0 && a <= 25.0:
		fmt.Printf("%s [0,25]\n", mensagem)
	case a > 25.0 && a <= 50.00:
		fmt.Printf("%s (25,50]\n", mensagem)
	case a > 50.0 && a <= 75.00:
		fmt.Printf("%s (50,75]\n", mensagem)
	case a > 75 && a <= 100.00:
		fmt.Printf("%s (75,100]\n", mensagem)
	default:
		fmt.Println("Fora de intervalo")
	}
}

func main() {
	var x float64
	fmt.Scanln(&x)
	calcularIntervalo(x)
}
