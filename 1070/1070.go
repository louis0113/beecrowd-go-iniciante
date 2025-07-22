package main

import (
	"fmt"
)

func main() {

	var x, i int

	fmt.Scanln(&x)
	if x%2 == 0 {
		i = x + 12
		loop(x, i)

	} else {
		i = x + 10
		loop(x, i)
	}

}

func loop(i, in int) {
	for c := i; c <= in; c++ {
		if c%2 == 1 {
			fmt.Println(c)
		}
	}
}
