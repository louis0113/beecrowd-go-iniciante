package main

import (
  "fmt"
)

const mensagem = "Intervalo"

func main() {

  var x float64

  fmt.Scanln(&x)

  switch{

    case x >= 0.0 && x <= 25.0: 
    fmt.Printf("%s [0,25]\n", mensagem)
    case x > 25.0 && x <= 50.00: 
    fmt.Printf("%s (25,50]\n", mensagem)
  case x > 50.0 && x <= 75.00:
    fmt.Printf("%s (50,75]\n", mensagem)
  case x > 75  && x <= 100.00:
    fmt.Printf("%s (75,100]\n", mensagem)
  default:
    fmt.Println("Fora de intervalo")
  } 


}
