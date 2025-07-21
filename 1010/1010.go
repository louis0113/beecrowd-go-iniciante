package main
 
import (
    "fmt"
)
 
func main() {

  var codP1, numP1, codP2, numP2 uint32
  
  var priceP1, priceP2 float32
  
  fmt.Scanf("%d %d %f\n", &codP1, &numP1, &priceP1)
  
  fmt.Scanf("%d %d %f\n", &codP2, &numP2, &priceP2)
  
  total := float32(numP1) * priceP1 + float32(numP2) * priceP2
  
  fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
  
}
