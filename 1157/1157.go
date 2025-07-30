package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Scanln(&n)

	for x := 1; x <= n; x++ {

		if n%x == 0 {
			fmt.Println(x)
		}

	}

}
