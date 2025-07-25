 
import (
    "fmt"
)
 
 const msg = "Novo grenal (1-sim 2-nao)"
 
func main() {

    var opt, inter, gre int
    
   
    
    tp, vi,vg, e := 0,0,0,0
   
    
    for{
        
        fmt.Scanf("%d %d\n", &inter, &gre)
        
        if inter > gre {
            vi++
            } else if gre > inter{
                vg++
                } else {
                    e++
                    }
                    
        fmt.Printf("%s\n", msg)
            
        fmt.Scanf("%d\n", &opt)
        
        tp++
        
        if opt == 2{
            break
            }
        }
    
    condicoes := [...]string {"Inter", "Gremio", "Empates"}
    
    valores := [...] int {vi,vg,e}
    
    fmt.Printf("%d grenais\n", tp)
    
    for x := 0; x < len(valores); x++{
            fmt.Printf("%s:%d\n", condicoes[x], valores[x])
        } 
        
        if vi > vg {
            fmt.Println("Inter venceu mais")
            } else if vg > vi{
                fmt.Println("Gremio venceu mais")
                } else {
                    fmt.Println("Nao houve vencedor")
                    }
    
    
}
