package main

import (
	"fmt"
)

func dentroIntervalo(v []int)(in, out int){
	for x := range len(v) {
		fmt.Scanln(&v[x])
		if v[x] >= 10 && v[x] <= 20 {
			in++
		} else {
			out++
		}
	}
	return
}

func main() {
	var input int
	fmt.Scanln(&input)
	values := make([]int, input)
	dentro, fora := dentroIntervalo(values)
	fmt.Printf("%d in\n%d out\n", dentro, fora)
}
