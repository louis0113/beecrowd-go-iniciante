package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanln(&x)
	initFunc(x)
}

func initFunc(a int){
	var b int
	if a%2 == 0 {
		b = a + 12
		loop(a, b)
	} else {
		b = a + 10
		loop(a, b)
	}

}
func loop(i, in int) {
	for c := i; c <= in; c++ {
		if c%2 == 1 {
			fmt.Println(c)
		}
	}
}
