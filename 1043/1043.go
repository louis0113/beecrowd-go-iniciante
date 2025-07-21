package main

import (
  "fmt"
)

func main() {

  var a,b,c float64

  fmt.Scanf("%g %g %g\n", &a, &b, &c)

  if a <  b + c && b < a + c && c < a + b{
    per := a + b + c
    fmt.Printf("Perimetro = %.1f\n", per)
  } else {

    traArea := ((a + b) * c) / 2

    fmt.Printf("Area = %.1f\n", traArea)
  }

}
