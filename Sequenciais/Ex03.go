package main

import "fmt"

func main() {
	var peso float64
	fmt.Print("Qual é o seu peso em kg? ")
	fmt.Scan(&peso)
	var altura float64
	fmt.Print("Qual é a sua altura em metros? ")
	fmt.Scan(&altura)
	var quadrado float64
	quadrado = altura * altura
	var imc float64
	imc = peso / quadrado
	fmt.Println("Seu IMC é: ", imc)
}
