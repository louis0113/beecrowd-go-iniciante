package main

import (
	"fmt"
)

func main() {

	var x, y int

	for {
		fmt.Scanf("%d %d\n", &x, &y)

		if x > y {
			fmt.Println("Decrescente")
		} else if x < y {
			fmt.Println("Crescente")
		} else {
			break
		}
	}

}
