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

            if i+j < 11 { 
                soma += values[i][j]
                cont++
            }
        }
    }

    if op == "S" {
        fmt.Printf("%.1f\n", soma)
    } else {
        fmt.Printf("%.1f\n", soma/float64(cont))
    }
}
