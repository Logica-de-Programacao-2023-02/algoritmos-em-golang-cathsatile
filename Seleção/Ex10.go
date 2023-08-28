package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Quantos anos você tem? ")
	fmt.Scan(&idade)
	if idade <= 9 {
		fmt.Print("Sua classificação é: Mirim")
	}
	if idade > 9 && idade <= 13 {
		fmt.Print("Sua classificação é: Infantil")
	}
	if idade > 13 && idade <= 17 {
		fmt.Print("Sua classificação é: Juvenil")
	}
	if idade >= 18 {
		fmt.Print("Sua classificação é: Adulto Maior")
	}
}
