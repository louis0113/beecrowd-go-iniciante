package main

import (
	"fmt"
)

func main() {
	var arr [10]int
	for x := range len(arr) {
		fmt.Scanln(&arr[x])
		if arr[x] <= 0 {
			arr[x] = 1
		}
		fmt.Printf("X[%d] = %d\n", x, arr[x])
	}
}
