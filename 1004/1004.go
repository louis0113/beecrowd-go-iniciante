package main
 
import (
    "fmt"
)
 
func main() {

    var x,y int
    
    fmt.Scanln(&x)
    
    fmt.Scanln(&y)
    
    prod := x * y
    
    fmt.Println("PROD =", prod)
    
}
