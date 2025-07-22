package main

import (
	"fmt"
)

func main() {

	for x := 1; x < 101; x++ {

		if x%2 == 0 {
			fmt.Println(x)
		}

	}

}
