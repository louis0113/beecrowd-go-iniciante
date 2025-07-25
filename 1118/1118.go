package main

import (
	"fmt"
)

const msg = "novo calculo (1-sim 2-nao)"

func main() {
	var nota float64
	var soma, media float64 = 0.0, 0.0
	var nv int = 0

	for {
		fmt.Scanln(&nota)

		if nota >= 0 && nota <= 10 {
			soma += nota
			nv++
		} else {
			fmt.Println("nota invalida")
		}

		if nv == 2 {
			media = soma / 2.0
			fmt.Printf("media = %.2f\n", media)

			for {
				fmt.Println(msg)
				var nn int
				fmt.Scanln(&nn)

				if nn == 1 {
					soma = 0.0
					nv = 0
					break
				} else if nn == 2 {
					return
				} else {

				}
			}
		}
	}
}
