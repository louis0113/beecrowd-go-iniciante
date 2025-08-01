package main

import (
	"fmt"
	"math"
)

func main() {

	var n, i, a int
	var is_prime bool

	fmt.Scanln(&n)

	for x := 0; x < n; x++ {

		is_prime = true
		fmt.Scanln(&i)

		value := math.Floor(math.Sqrt(float64(i)))

		a = 2

		for a <= int(value) {

			if i%a == 0 {
				is_prime = false
				break
			}
			a++
		}

		if is_prime {
			fmt.Println(i, "eh primo")
		} else {
			fmt.Println(i, "nao eh primo")
		}

	}
}
