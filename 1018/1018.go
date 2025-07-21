package main

import (
  "fmt"
)

const mensagem = "nota(s) de R$"

func main() {

  var n int

  fmt.Scan(&n)

  notas := [...]int{100,50,20,10,5,2,1}

  fmt.Println(n)

  for x:= 0; x < len(notas); x++ {
    total := n / notas[x]

    n %= notas[x]

    fmt.Printf("%v %s %v,00\n", total, mensagem, notas[x])

  }

}
