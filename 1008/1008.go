package main

import (
	"fmt"
)

func main() {

	var numFunc, hour int
	var salary float64

	fmt.Scanf("%d\n%d\n%g\n", &numFunc, &hour, &salary)

	totalSalary := float64(hour) * salary

	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", numFunc, totalSalary)

}
