package main

import (
	"fmt"
)

func condicoes(v, x, y, z int) {
	if x > y && z > v && (y+z) > (v+x) && y > 0 && z > 0 && v%2 == 0 {
		fmt.Println("Valores aceitos")
	} else {
		fmt.Println("Valores nao aceitos")
	}
}

func main() {
	var a, b, c, d int
	fmt.Scanf("%d %d %d %d\n", &a, &b, &c, &d)

	condicoes(a,b,c,d)
}
