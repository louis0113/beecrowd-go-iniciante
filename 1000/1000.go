package main

import (
	"fmt"
)

func helloWorld() string {
	var hw string
	hw = "Hello World"
	return hw
}

func main() {
	hello := helloWorld()
	fmt.Println(hello)
}
