package main

import (
	"fmt"
)

const msg = "Total de"
const msg2 = "Percentual de"

func main() {
	var n int
	fmt.Scanln(&n)
	quan := make([]int, n)
	tipo := make([]string, n)
	total, coelhos, ratos, sapos := 0, 0, 0, 0
	perc_c, perc_r, perc_s := 0.0, 0.0, 0.0
	for i := range n {
		fmt.Scanf("%d %s\n", &quan[i], &tipo[i])
		switch tipo[i]{
			case "C":
				coelhos += quan[i]
			case "R":
				ratos += quan[i]
			default:
				sapos += quan[i]
		}
		total += quan[i]
	}
	perc_c = (float64(coelhos) * 100) / float64(total)
	perc_r = (float64(ratos) * 100) / float64(total)
	perc_s = (float64(sapos) * 100) / float64(total)
	quantidades := [...]int{coelhos, ratos, sapos}
	porcentagens := [...]float64{perc_c, perc_r, perc_s}
	animais := [...]string{"coelhos", "ratos", "sapos"}
	fmt.Printf("Total: %d cobaias\n", total)
	for a := range 3 {
		fmt.Printf("%s %s: %d\n", msg, animais[a], quantidades[a])
	}
	for x := range 3 {
		fmt.Printf("%s %s: %.2f %%\n", msg2, animais[x], porcentagens[x])
	}
}
