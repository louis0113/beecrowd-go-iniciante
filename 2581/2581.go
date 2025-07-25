package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Scanln(&n)

	perguntas := make([]string, n)

	for x := 0; x < n; x++ {

		fmt.Scanln(&perguntas[x])

		fmt.Println("I am Toorg!")
	}

}
