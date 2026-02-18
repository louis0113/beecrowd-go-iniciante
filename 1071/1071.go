package main

import (
	"fmt"
)

func soma(a,b int)(soma int){
	soma = 0
	for c := (b + 1); c < a; c++ {
		if c%2 == 1 || c%2 == -1 {
			soma += c
		}
	}
return
}

func main() {
	var x, y int
	fmt.Scanf("%d\n%d\n",&x, &y)
	s := soma(x,y)
	fmt.Println(s)
}
