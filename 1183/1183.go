package main

import (
	"fmt"
)

func main() {
	var (
		M [12][12]float64
		O string
		r float64
		t int
	)

	fmt.Scanln(&O)

	for x := 0; x < 12; x++ {
		for y := 0; y < 12; y++ {
			fmt.Scan(&M[x][y])
		}
	}

	for linha := 0; linha < 12; linha++ {
		for coluna := linha + 1; coluna < 12; coluna++ {
			r += M[linha][coluna]
			t++
		}
	}

	if O == "S" {
		fmt.Printf("%.1f\n", r)
	}

	if O == "M" {
		if t > 0 {
			r /= float64(t)
		}
		fmt.Printf("%.1f\n", r)
	}
}
