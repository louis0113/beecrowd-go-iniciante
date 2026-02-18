package main

import (
	"fmt"
)

func numerosPares(nums [5]int)(c int){
	c = 0;
	for x := range 5{
		fmt.Scanln(&nums[x])
		if nums[x]%2 == 0 {
			c++
		}
	}
	return
}
func main() {
	var values [5]int
	c := numerosPares(values)
	fmt.Printf("%d valores pares\n", c)
}
