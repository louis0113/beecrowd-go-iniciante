package main

import (
	"fmt"
	"unicode"
)

func main() {

	var (
		n, n1, n2, r int
		v            []rune
		v2           [][]rune
	)

	fmt.Scan(&n)

	arr := make([]string, n)

	for x := range len(arr) {
		fmt.Scan(&arr[x])
		v = []rune(arr[x])
		v2 = append(v2, v)
	}

	for y := range n {
		n1, n2 = int(v2[y][0] - '0'), int(v2[y][2] - '0')
		if n1 == n2 {
			r = n1 * n2
		} else if unicode.IsUpper(v2[y][1]) {
			r = n2 - n1
		} else if unicode.IsLower(v2[y][1]) {
			r = n1 + n2
		}

		fmt.Println(r)

	}
}
