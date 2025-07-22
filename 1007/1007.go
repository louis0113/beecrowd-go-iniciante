package main

import (
	"fmt"
)

func main() {

	var a, b, c, d int

	fmt.Scanf("%d\n%d\n%d\n%d\n", &a, &b, &c, &d)

	diff := a*b - c*d

	fmt.Println("DIFERENCA =", diff)

}
