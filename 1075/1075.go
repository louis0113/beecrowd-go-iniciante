package main

import (
	"fmt"
)
func moduleTwo(a int){
	for x := 1; x <= 10000; x++ {
		if x%a == 2 {
			fmt.Println(x)
		}
	}
}
func main() {
	var n int
	fmt.Scanln(&n)
	moduleTwo(n)
}
