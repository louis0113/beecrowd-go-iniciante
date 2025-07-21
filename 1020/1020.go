package main

import (
  "fmt"
)
const anoDia = 365
const mesDia = 30
func main() {

  var dias, anos, meses int

  fmt.Scanln(&dias)

  anos = dias / anoDia
  dias%= anoDia

  meses = dias / mesDia
  dias%= mesDia

  fmt.Printf("%v ano(s)\n%v mes(es)\n%v dia(s)\n", anos, meses, dias)

}
