package main

import (
  "fmt"
  "math"
  "sort"
)

func main() {
  var a, b, c float64

  fmt.Scanf("%g %g %g\n", &a, &b, &c)

  lados := []float64{a, b, c}

  // Ordena os lados em ordem decrescente
  sort.Slice(lados, func(i, j int) bool {
    return lados[i] > lados[j]
  })

  // Verifica se é um triângulo válido
  switch {
  case lados[0] >= lados[1] + lados[2]:
    fmt.Println("NAO FORMA TRIANGULO")

  case lados[0] < lados[1] + lados[2]:
    // Verifica o tipo de triângulo
    switch {
    case math.Pow(lados[0], 2) == math.Pow(lados[1], 2) + math.Pow(lados[2], 2):
      fmt.Println("TRIANGULO RETANGULO")
    case math.Pow(lados[0], 2) > math.Pow(lados[1], 2) + math.Pow(lados[2], 2):
      fmt.Println("TRIANGULO OBTUSANGULO")
    case math.Pow(lados[0], 2) < math.Pow(lados[1], 2) + math.Pow(lados[2], 2):
      fmt.Println("TRIANGULO ACUTANGULO")
    }

    // Verifica o tipo de triângulo baseado nos lados
    switch {
    case lados[0] == lados[1] && lados[1] == lados[2]:
      fmt.Println("TRIANGULO EQUILATERO")
    case lados[0] == lados[1] || lados[0] == lados[2] || lados[1] == lados[2]:
      fmt.Println("TRIANGULO ISOSCELES")
    }
  }
}
