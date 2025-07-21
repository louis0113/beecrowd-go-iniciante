package main

import (
  "fmt"
  "math"
)

func main() {

  var a,b,c uint32

  fmt.Scanf("%v %v %v\n", &a, &b , &c)

  res1 := maior(float64(a),float64(b))

  resf := maior(float64(res1), float64(c))

  fmt.Printf("%v eh o maior\n", resf)

}

func maior (x,y float64) uint32 {

  maior := (x + y + math.Abs(x-y)) / 2

  r := uint32(maior)

  return r

}
