package main

import (
	"fmt"
)

const m1 = "O JOGO DUROU"
const m2 = "HORA(S)"
const horaMax uint8 = 24

func main() {
	var hi, hf, hora uint8
	fmt.Scanf("%d %d\n", &hi, &hf)
	if hi > hf {
		hora = calcularHora1(hi,hf)
		mostrarHoras(m1, hora, m2)
	} else if hf > hi {
		hora = calcularHora2(hi,hf)
		mostrarHoras(m1, hora, m2)
	} else {
		mostrarHoras(m1, horaMax, m2)
	}
}

func calcularHora1(horaInicial, horaFinal uint8) uint8{
		sobra := horaMax - horaInicial
		hora := sobra + horaFinal
		return hora
}

func calcularHora2(horaInicial, horaFinal uint8) uint8{
	hora := horaFinal - horaInicial
	return hora
}

func mostrarHoras(men1 string, x uint8, men2 string) {
	fmt.Printf("%s %d %s\n", men1, x, men2)
}
