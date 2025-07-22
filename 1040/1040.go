package main

import (
	"fmt"
)

const peso1, peso2, peso3, peso4 = 2.0, 3.0, 4.0, 1.0

func main() {

	var n1, n2, n3, n4 float32

	fmt.Scanf("%f %f %f %f\n", &n1, &n2, &n3, &n4)

	m := media(n1, n2, n3, n4)
	fmt.Printf("Media: %.1f\n", m)

	if m >= 7.0 {
		fmt.Println("Aluno aprovado.")
	} else if m >= 5.0 && m < 7.0 {
		fmt.Println("Aluno em exame.")
		var nf, mf float32
		fmt.Scanln(&nf)
		mf = mediaFinal(m, nf)

		fmt.Printf("Nota do exame: %.1f\n", nf)

		if mf >= 5.0 {
			fmt.Println("Aluno aprovado.")
		} else {
			fmt.Println("Aluno reprovado.")
		}

		fmt.Printf("Media final: %.1f\n", mf)
	} else {
		fmt.Println("Aluno reprovado.")
	}

}

func media(a, b, c, d float32) float32 {

	totalPesos := peso1 + peso2 + peso3 + peso4

	media := (a*float32(peso1) + b*float32(peso2) + c*float32(peso3) + d*float32(peso4)) / float32(totalPesos)

	return media

}

func mediaFinal(x, y float32) float32 {
	mediaf := (x + y) / 2

	return mediaf
}
