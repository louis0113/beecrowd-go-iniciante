package main
 
import (
    "fmt"
)
const pesoA, pesoB, pesoC = 2.0,3.0,5.0
 
func main() {

var a,b,c float64

fmt.Scanf("%g\n%g\n%g\n", &a, &b, &c)


totalPesos := pesoA + pesoB + pesoC

media := (pesoA * a + pesoB * b + pesoC * c ) / totalPesos

fmt.Printf("MEDIA = %.1f\n", media)
}
