package main

import (
	"fmt"
)

func main() {

	var x, y int

	for {

		fmt.Scanf("%d %d\n", &x, &y)

		if x > 0 && y > 0 {
			fmt.Println("primeiro")
		} else if x < 0 && y > 0 {
			fmt.Println("segundo")
		} else if x < 0 && y < 0 {
			fmt.Println("terceiro")
		} else if x > 0 && y < 0 {
			fmt.Println("quarto")
		} else {
			break
		}

	}

}
