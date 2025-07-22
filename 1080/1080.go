package main

import (
	"fmt"
)

func main() {

	var values [100]int
	maior, p, pm := 0, 0, 0
	for i := 0; i < 100; i++ {

		fmt.Scanln(&values[i])

		if values[i] > maior {
			maior = values[i]
			pm = p + 1
		}
		p++
	}

	fmt.Println(maior)
	fmt.Println(pm)
}
