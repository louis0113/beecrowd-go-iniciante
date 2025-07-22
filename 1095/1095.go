package main

import (
	"fmt"
)

func main() {

	i, j := 1, 60

	for {

		if j == 0 {
			fmt.Printf("I=%d J=%d\n", i, j)
			break
		} else {
			fmt.Printf("I=%d J=%d\n", i, j)
		}

		i += 3
		j -= 5

	}

}
