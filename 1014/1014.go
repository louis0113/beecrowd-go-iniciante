package main
 
import (
    "fmt"
)
 
func main() {

var x int
var y float32

fmt.Scanf("%d\n%f\n ", &x, &y)


resultado := float32(x) / y

fmt.Printf("%.3f km/l\n", resultado)

}
