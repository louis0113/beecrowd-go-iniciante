package main

import (
	"fmt"
)

func imparRange(a int){
	for c := range a {
		if c%2 == 1 {
			fmt.Println(c)
		}
	}
}

func main() {
	var x int
	fmt.Scanln(&x)
	imparRange(x)
}
