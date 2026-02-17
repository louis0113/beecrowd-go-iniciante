package main

import (
	"fmt"
)

const anoDia = 365
const mesDia = 30

func calcularTempo(d int) (anos, meses, dias int) {
	anos = d / anoDia
	d %= anoDia
	meses = d / mesDia
	dias = d % mesDia
	return
}

func main() {
	var dias, anos, meses int
	fmt.Scanln(&dias)
	anos, meses, dias = calcularTempo(dias)
	fmt.Printf("%v ano(s)\n%v mes(es)\n%v dia(s)\n", anos, meses, dias)
}
