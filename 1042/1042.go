package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d\n", &a, &b, &c)
	numbers(a,b,c)
}

func numbers(x,y,z int){
	numeros := []int{x, y, z}
	numerosCresc := []int{x, y, z}
	sort.Ints(numerosCresc)
	printNumbers(numerosCresc)
	fmt.Println()
	printNumbers(numeros)

}

func printNumbers(nums []int) {
	for i := range len(nums) {
		fmt.Println(nums[i])
	}
}
