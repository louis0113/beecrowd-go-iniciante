package main

import "fmt"

func qualAnimal(vertebra, alimentacao, classe string) (animal string) {

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
        return
}

func main() {
        var vertebra, classe, alimentacao string
        fmt.Scanf("%s\n%s\n%s\n", &vertebra, &classe, &alimentacao)
        animal := qualAnimal(vertebra, alimentacao, classe)
        fmt.Println(animal)
}
