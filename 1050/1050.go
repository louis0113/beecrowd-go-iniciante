package main

import (
	"fmt"
)

 func qualEstado(ddd uint8)(estado string){
	switch ddd {
	case 61:
		estado = "Brasilia"
	case 71:
		estado = "Salvador"
	case 11:
		estado = "Sao Paulo"
	case 21:
		estado = "Rio de Janeiro"
	case 32:
		estado = "Juiz de Fora"
	case 19:
		estado = "Campinas"
	case 27:
		estado = "Vitoria"
	case 31:
		estado = "Belo Horizonte"
	default:
		estado = "DDD nao cadastrado"
	}
	return
 }
func main() {
	var ddd uint8
	fmt.Scanln(&ddd)
	estado := qualEstado(ddd)
	fmt.Println(estado)
}
