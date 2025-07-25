package main

import (
	"fmt"
)

func main() {

	var nota float64
	cn := 0

	soma := 0.0
	for {

		fmt.Scanln(&nota)

		if nota >= 0.0 && nota <= 10.0 {
			soma += nota
			cn++
		} else {
			fmt.Println("nota invalida")
		}

		if cn == 2 {
			break
		}
	}

	r := soma / float64(cn)
	fmt.Printf("media = %.2f\n", r)
}
