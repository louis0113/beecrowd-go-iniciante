package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const minutosSegundos uint = 60
const horaSegundos uint = minutosSegundos * 60
const diaSegundos uint = 24 * horaSegundos

func input(a, b string) (x, y string) {
	reader := bufio.NewReader(os.Stdin)
	a, _ = reader.ReadString('\n')
	b, _ = reader.ReadString('\n')
	x, y = a, b
	return
}

func getValues(a, b string) (day, hour, minutes, seconds uint) {
	str := strings.Fields(a)
	hms := strings.Fields(b)
	new_hms := ""
	for i := range len(hms) {
		new_hms = new_hms + hms[i]
	}
	arr_hms := strings.Split(new_hms, ":")
	day, hour, minutes, seconds = stringToUint(str[1]), stringToUint(arr_hms[0]), stringToUint(arr_hms[1]), stringToUint(arr_hms[2])
	return
}

func stringToUint(str string) uint {
	num, _ := strconv.ParseInt(str, 10, 64)
	return uint(num)
}

func calculardiasHorasMinutosSegundos(arr_inicial, arr_final [4]uint) (arr_atual [4]uint) {
	var totalnicial, totalFinal uint = 0, 0
	segundosConvertidos := [3]uint{diaSegundos, horaSegundos, minutosSegundos}
	for x := range 4 {
		if x != 3 {
			totalnicial += arr_inicial[x] * segundosConvertidos[x]
			totalFinal += arr_final[x] * segundosConvertidos[x]
		} else {
			totalnicial += arr_inicial[x]
			totalFinal += arr_final[x]
		}
	}
	var totalAtual uint = totalFinal - totalnicial
	resto := totalAtual
	for y := range 4 {
		if y != 3 {
			arr_atual[y] = resto / segundosConvertidos[y]
			resto %= segundosConvertidos[y]
		} else {
			arr_atual[y] = resto
		}
	}
	return
}

func showData(data [4]uint, mensagens [4]string) {
	for x := range len(data) {
		fmt.Printf("%d %s\n", data[x], mensagens[x])
	}
}

func main() {
	msgs := [4]string{"dia(s)", "hora(s)", "minuto(s)", "segundo(s)"}
	var a1, b1, a2, b2 string
	var d1, h1, m1, s1, d2, h2, m2, s2 uint
	a1, b1 = input(a1, b1)
	a2, b2 = input(a2, b2)
	d1, h1, m1, s1 = getValues(a1, b1)
	d2, h2, m2, s2 = getValues(a2, b2)
	data1 := [4]uint{d1, h1, m1, s1}
	data2 := [4]uint{d2, h2, m2, s2}
	var finalData [4]uint = calculardiasHorasMinutosSegundos(data1, data2)
	showData(finalData, msgs)
}
