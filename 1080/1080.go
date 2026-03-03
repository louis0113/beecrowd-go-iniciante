package main

import (
	"fmt"
)
func calculos(){

}

func main() {
	var values [100]int
	maior, p, pm := 0, 0, 0
	for i := range 100 {
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
