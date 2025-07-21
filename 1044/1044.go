package main

import (
  "fmt"
)

func main() {

  var a, b int

  fmt.Scanf("%v %v\n", &a, &b)

  if b % a == 0 || a % b == 0 {
    fmt.Println("Sao Multiplos")
  } else {
    fmt.Println("Nao sao Multiplos")
  }

}

