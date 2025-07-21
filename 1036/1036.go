package main

import (
  "fmt"
  "math"
)

func main() {

  var a,b,c float64

  fmt.Scanf("%g %g %g\n", &a, &b, &c)

  d := delta(a,b,c)

  if d <= 0 || a == 0 {
    fmt.Println("Impossivel calcular")
  } else {
    x1,x2 := bhaskara(a,b,d)

    fmt.Printf("R1 = %.5f\nR2 = %.5f\n", x1,x2)
  }


}


func delta(x,y,z float64) float64{

  delta := math.Pow(y,2) - (4 * x * z) 

  return delta

}


func bhaskara(x,y, delta float64) (x1,x2 float64){


  x1 = ((-1 * y) + math.Sqrt(delta)) / (2 * x)    
  x2 = ((-1 * y) - math.Sqrt(delta)) / (2 * x)

  return 
}
