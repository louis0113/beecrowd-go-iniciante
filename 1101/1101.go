package main

import (
	"fmt"
)

const msg = "Sum="

func main() {

	var x, y, sum int
	var m string

	for {

		fmt.Scanf("%d %d\n", &x, &y)

		if x <= 0 || y <= 0 {
			break
		} else {
			sum = 0
			m = ""

			min := x
			max := y

			if y < x {
				min = y
				max = x
			}

			for c := min; c <= max; c++ {

				m += fmt.Sprintf("%d ", c)
				sum += c
			}

			fmt.Printf("%s%s%d\n", m, msg, sum)

		}

	}

}
