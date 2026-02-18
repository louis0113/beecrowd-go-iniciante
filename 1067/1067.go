package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanln(&x)
	for c := range x {
		if c%2 == 1 {
			fmt.Println(c)
		}
	}
}
