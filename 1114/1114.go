package main

import (
	"fmt"
)

func main() {

	var senha string

	for {

		fmt.Scanln(&senha)

		if senha == `2002` {
			fmt.Println("Acesso Permitido")
			break
		} else {
			fmt.Println("Senha Invalida")
		}

	}
}
