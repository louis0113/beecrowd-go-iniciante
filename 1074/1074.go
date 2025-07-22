package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Scanln(&n)

	values := make([]int, n)

	for x := 0; x < len(values); x++ {

		fmt.Scanln(&values[x])

		if values[x] == 0 {
			fmt.Println("NULL")
		} else if values[x] > 0 && values[x]%2 == 0 {
			fmt.Println("EVEN POSITIVE")
		} else if values[x] < 0 && values[x]%2 == 0 {
			fmt.Println("EVEN NEGATIVE")

		} else if values[x] > 0 && values[x]%2 != 0 {
			fmt.Println("ODD POSITIVE")
		} else {
			fmt.Println("ODD NEGATIVE")
		}
	}
}
