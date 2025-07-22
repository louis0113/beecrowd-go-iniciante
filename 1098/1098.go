package main

import (
	"fmt"
)

func main() {

	i, j := 0.0, 1.0

	for {
		if i >= 2.0 {
			for x := 0.0; x <= 2.0; x += 1.0 {
				fmt.Printf("I=%d J=%d\n", int(i), int(x+j))
			}
			break
		} else if i == 0.0 || i == 1.0 {
			for x := 0.0; x <= 2.0; x += 1.0 {
				fmt.Printf("I=%d J=%d\n", int(i), int(x+j+0.0001))
			}

		} else {
			if i <= 1.8 {
				for x := 0.0; x <= 2.0; x += 1.0 {
					fmt.Printf("I=%.1f J=%.1f\n", i, (x + j))
				}
			}

		}

		i += 0.2
		j += 0.2

	}
}
