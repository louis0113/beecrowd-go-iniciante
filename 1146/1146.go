package main

import (
	"fmt"
)

func main() {

	var n int

	for {

		fmt.Scanln(&n)

		if n == 0 {
			break
		}

		for x := 1; x <= n; x++ {

			if x == n {
				fmt.Printf("%d\n", x)
			} else {
				fmt.Printf("%d ", x)
			}
		}

	}
}
