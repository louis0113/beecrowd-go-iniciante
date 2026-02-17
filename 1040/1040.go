package main

import (
	"fmt"
)

const peso1, peso2, peso3, peso4 float32 = 2.0, 3.0, 4.0, 1.0

func main() {
	var n1, n2, n3, n4 float32
	fmt.Scanf("%f %f %f %f\n", &n1, &n2, &n3, &n4)
	mensagemStatus(n1, n2, n3, n4)
}

func mensagemStatus(nota1, nota2, nota3, nota4 float32) {
	m := media(nota1, nota2, nota3, nota4)
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
	media := (a*peso1 + b*peso2 + c*peso3 + d*peso4) / totalPesos
	return media
}

func mediaFinal(x, y float32) float32 {
	mediaf := (x + y) / 2
	return mediaf
}
