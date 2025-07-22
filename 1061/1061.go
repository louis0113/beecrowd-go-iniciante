package main

import (
	"fmt"
)

func main() {
	var di, hi, mi, si, df, hf, mf, sf int

	// Passando os endereços das variáveis (ponteiros) para a função input
	input(&di, &hi, &mi, &si)
	input(&df, &hf, &mf, &sf)

	totalSegInicial := conversaoSeg(di, hi, mi, si)
	totalSegFinal := conversaoSeg(df, hf, mf, sf)

	totalSegAtual := totalSegFinal - totalSegInicial

	// Se o tempo final for menor que o inicial, significa que o evento passou para o dia seguinte
	if totalSegAtual < 0 {
		totalSegAtual += (24 * 3600) + (48 * 3600) + 60*60 + 60
	}

	out(totalSegAtual)
}

func input(dia, hora, min, seg *int) {
	fmt.Scanf("Dia %d\n", dia)
	fmt.Scanf("%d : %d : %d\n", hora, min, seg)
}

func conversaoSeg(dia, hora, min, seg int) int {
	diaSeg := dia * 24 * 3600
	horaSeg := hora * 3600
	minSeg := min * 60

	segundos := diaSeg + horaSeg + minSeg + seg
	return segundos
}

func out(segundos int) {
	dias := segundos / (24 * 3600)
	segundos %= (24 * 3600)
	horas := segundos / 3600
	segundos %= 3600
	minutos := segundos / 60
	segundos %= 60

	tempos := [4]int{dias, horas, minutos, segundos}
	mess := [4]string{"dia(s)", "hora(s)", "minuto(s)", "segundo(s)"}

	for x := 0; x < len(tempos); x++ {
		fmt.Printf("%d %s\n", tempos[x], mess[x])
	}
}
