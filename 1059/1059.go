package main

import (
	"fmt"
)

func mostrarPares() {
	for x := range 101{
		if x%2 == 0 && x != 0 {
			fmt.Println(x)
		}
	}
}

func main() {
	mostrarPares()
}
