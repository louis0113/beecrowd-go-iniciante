package main

import (
	"fmt"
)

const m1 = "O JOGO DUROU"
const m2 = "HORA(S)"
const horaMax = 24

func main() {

	var hi, hf, hora uint8

	fmt.Scanf("%d %d\n", &hi, &hf)

	if hi > hf {
		sobra := uint8(horaMax) - hi
		hora = sobra + hf
		horas(m1, hora, m2)

	} else if hf > hi {
		hora = hf - hi
		horas(m1, hora, m2)

	} else {
		horas(m1, uint8(horaMax), m2)
	}
}

func horas(men1 string, x uint8, men2 string) {

	fmt.Printf("%s %d %s\n", men1, x, men2)
}
