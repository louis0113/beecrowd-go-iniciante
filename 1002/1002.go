package main
 
import (
    "fmt"
    "math"
)
 const Pi = 3.14159
 
func main() {

    var r float64
    
    fmt.Scanln(&r)
    
    area := Pi * math.Pow(r,2)
    
    fmt.Printf("A=%.4f\n", area)

}
