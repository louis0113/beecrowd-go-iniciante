package main

import (
	"fmt"
	"math"
)

const msg = "^2 ="

func pow(a int){
	pot := 0.0
	for x := range a{
		if x%2 == 0 {
			pot = math.Pow(float64(x), 2)
			fmt.Printf("%d%s %d\n", x, msg, int(pot))
		}
	}
}
func main() {
	var n int
	fmt.Scanln(&n)
	pow(n)
}
