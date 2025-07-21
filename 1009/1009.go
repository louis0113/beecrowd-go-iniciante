package main
 
import (
    "fmt"
)
 const comissao = 15
 
func main() {

    var name string
    var salario, vendas float64
    
    fmt.Scanf("%s\n%g\n%g\n", &name, &salario, &vendas)
    
    var bonus float64 =  vendas * comissao / 100
    novo_salario :=  salario + bonus
    
    fmt.Printf("TOTAL = R$ %.2f\n", novo_salario)
    
}
