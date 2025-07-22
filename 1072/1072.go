package main

import (
	"fmt"
)

func main() {

	var input int
	in, out := 0, 0
	fmt.Scanln(&input)

	values := make([]int, input)

	for x := 0; x < len(values); x++ {
		fmt.Scanln(&values[x])
		if values[x] >= 10 && values[x] <= 20 {
			in++
		} else {
			out++
		}

	}

	fmt.Printf("%d in\n%d out\n", in, out)

}
