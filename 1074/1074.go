package main

import (
	"fmt"
)
func evenOrOdd(v []int){
	for x := range len(v) {
		fmt.Scanln(&v[x])
		if v[x] == 0 {
			fmt.Println("NULL")
		} else if v[x] > 0 && v[x]%2 == 0 {
			fmt.Println("EVEN POSITIVE")
		} else if v[x] < 0 && v[x]%2 == 0 {
			fmt.Println("EVEN NEGATIVE")

		} else if v[x] > 0 && v[x]%2 != 0 {
			fmt.Println("ODD POSITIVE")
		} else {
			fmt.Println("ODD NEGATIVE")
		}
	}

}
func main() {
	var n int
	fmt.Scanln(&n)
	values := make([]int, n)
	evenOrOdd(values)
}
