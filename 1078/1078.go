package main

import (
	"fmt"
)

func mult(a int){
	var mult int
	for x := 1; x < 11; x++ {
		mult = x * a
		fmt.Printf("%d x %d = %d\n", x, a, mult)
	}

}
func main() {
	var n int
	fmt.Scanln(&n)
	mult(n)
}
