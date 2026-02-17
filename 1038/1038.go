package main

import (
	"fmt"
)

const mensagem = "Total: R$"

func main() {
	var codProduto, quanProduto int
	produtos := [5]float64{4.00, 4.50, 5.00, 2.00, 1.50}
	fmt.Scanf("%v %v\n", &codProduto, &quanProduto)

	codigoProduto(quanProduto, codProduto, produtos)
}

func codigoProduto(quan, cod int, prod [5]float64){
	switch cod {
	case 1:
		calcularTotal(quan, prod[0])
	case 2:
		calcularTotal(quan, prod[1])
	case 3:
		calcularTotal(quan, prod[2])
	case 4:
		calcularTotal(quan, prod[3])
	case 5:
		calcularTotal(quan, prod[4])
	}
}

func calcularTotal(quan int, preco float64) {
	total := float64(quan) * preco
	fmt.Printf("%s %.2f\n", mensagem, total)
}
