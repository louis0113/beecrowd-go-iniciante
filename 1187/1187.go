package main

import "fmt"

func main() {

    var (
        op     string
        values [12][12]float64
        soma   float64
        cont   int
    )

    fmt.Scan(&op)

    for i := range 12 {
        for j := range 12 {
            fmt.Scan(&values[i][j])

			if i > 0 && i < 11 && j < 5{
				soma += 
			}
        }
    }

    if op == "S" {
        fmt.Printf("%.1f\n", soma)
    } else {
        fmt.Printf("%.1f\n", soma/float64(cont))
    }
}
