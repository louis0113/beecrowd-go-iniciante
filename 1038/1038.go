package main

import (
  "fmt"
)

const mensagem = "Total: R$" 

func main() {

  var codProduto, quanProduto int

  produtos := [...]float64{4.00,4.50,5.00,2.00,1.50}
  fmt.Scanf("%v %v\n", &codProduto, &quanProduto)

  switch codProduto{

  case 1:
    calcularTotal(quanProduto, produtos[0])
  case 2:
    calcularTotal(quanProduto, produtos[1])
  case 3:
    calcularTotal(quanProduto, produtos[2])
  case 4:
    calcularTotal(quanProduto, produtos[3])
  case 5:
    calcularTotal(quanProduto, produtos[4])

  }

}

func calcularTotal (quan int , preco float64){
  total := float64(quan) * preco

  fmt.Printf("%s %.2f\n", mensagem, total)
}
