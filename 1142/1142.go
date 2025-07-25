package main

import (
	"fmt"
)

const msg = "PUM"

func main() {

	var p int
	var m string
	c := 0

	fmt.Scanln(&p)

	for x := 0; x < p; x++ {
		m = ""
		for i := 0; i < 3; i++ {
			c++
			m += fmt.Sprintf("%d ", c)
		}
		fmt.Printf("%s%s\n", m, msg)
		c++
	}
}
