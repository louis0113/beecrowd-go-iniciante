package main

import "fmt"

func main() {
	var vertebra, classe, alimentacao string
	var animal string

	fmt.Scanln(&vertebra)
	fmt.Scanln(&classe)
	fmt.Scanln(&alimentacao)

	switch vertebra {
	case "vertebrado":
		switch classe {
		case "ave":
			switch alimentacao {
			case "carnivoro":
				animal = "aguia"
			case "onivoro":
				animal = "pomba"
			}
		case "mamifero":
			switch alimentacao {
			case "onivoro":
				animal = "homem"
			case "herbivoro":
				animal = "vaca"
			}
		}
	case "invertebrado":
		switch classe {
		case "inseto":
			switch alimentacao {
			case "hematofago":
				animal = "pulga"
			case "herbivoro":
				animal = "lagarta"
			}
		case "anelideo":
			switch alimentacao {
			case "hematofago":
				animal = "sanguessuga"
			case "onivoro":
				animal = "minhoca"
			}
		}
	}
	
	fmt.Println(animal)
}
