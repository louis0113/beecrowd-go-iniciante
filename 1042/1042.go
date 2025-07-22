package main

import (
	"fmt"
	"sort"
)

func main() {

	var a, b, c int

	fmt.Scanf("%d %d %d\n", &a, &b, &c)

	numeros := []int{a, b, c}

	numerosCresc := []int{a, b, c}

	sort.Ints(numerosCresc)

	print(numerosCresc)
	fmt.Println()
	print(numeros)

}

func print(nums []int) {

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])

	}

}
