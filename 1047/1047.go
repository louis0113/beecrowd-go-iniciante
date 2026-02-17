package main

import "fmt"

const maxMin = 60
const maxHour = 24
const m3 = "MINUTO(S)"
const m2 = "HORA(S)"
const m1 = "O JOGO DUROU"

func horasMinutos(horaInicial, minutoInicial, horaFinal, minutoFinal int) (horaAtual, minutoAtual int) {

	if horaInicial < horaFinal && minutoInicial == minutoFinal {
		horaAtual = horaFinal - horaInicial
		minutoAtual = 0
	} else if horaInicial > horaFinal && minutoInicial == minutoFinal {
		horaAtual = (maxHour - horaInicial) + horaFinal
		minutoAtual = 0
	} else if horaInicial == horaFinal && minutoInicial < minutoFinal {
		horaAtual = 0
		minutoAtual = minutoFinal - minutoInicial
	} else if horaInicial == horaFinal && minutoInicial > minutoFinal {
		horaAtual = 0
		minutoAtual = (maxMin - minutoInicial) + minutoFinal
	} else if horaInicial < horaFinal && minutoInicial < minutoFinal {
		horaAtual = horaFinal - horaInicial
		minutoAtual = minutoFinal - minutoInicial
	} else if horaInicial > horaFinal && minutoInicial > minutoFinal {
		horaAtual = (maxHour - horaInicial) + horaFinal
		minutoAtual = (maxMin - minutoInicial) + minutoFinal
	} else if horaInicial < horaFinal && minutoInicial > minutoFinal && (horaFinal-horaInicial) == 1 {
		horaAtual = (horaFinal - horaInicial) - 1
		minutoAtual = (maxMin - minutoInicial) + minutoFinal
	} else if horaInicial > horaFinal && minutoInicial < minutoFinal {
		horaAtual = (maxHour - horaInicial) + horaFinal
		minutoAtual = minutoFinal - minutoInicial
	} else if horaInicial == horaFinal && minutoInicial == horaFinal {
		horaAtual = 24
		minutoAtual = 0
	} else if horaInicial < horaFinal && minutoInicial > minutoFinal && (horaFinal-horaInicial) != 1 {
		horaAtual = (horaFinal - horaInicial) - 1
		minutoAtual = (maxMin - minutoInicial) + minutoFinal
	}
	return
}

func main() {
	var hi, mi, hf, mf, ha, ma int
	fmt.Scanf("%d %d %d %d\n", &hi, &mi, &hf, &mf)
	ha, ma = horasMinutos(hi, mi, hf, mf)
	fmt.Printf("%s %d %s E %d %s\n", m1, ha, m2, ma, m3)
}
