package main

import (
	"fmt"
)

func totalPositivo(arr [6]float64) int {
	c := 0
	for x := range 6 {
		fmt.Scanln(&arr[x])
		if arr[x] > 0.0 {
			c++
		}
	}
	return c
}

func main() {
	var values [6]float64
	r := totalPositivo(values)
	fmt.Printf("%d valores positivos\n", r)
}
