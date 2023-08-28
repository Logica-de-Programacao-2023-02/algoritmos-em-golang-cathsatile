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
	fmt.Printf("Seu IMC é: %.2f ", imc)
	fmt.Println()

	if imc < 18.5 {
		fmt.Println("Você está abaixo do peso ideal! ")
	}
	if imc >= 18.5 && imc <= 25.0 {
		fmt.Println("Você está dentro do peso ideal! ")
	}
	if imc > 25.0 {
		fmt.Println("Você está acima do peso ideal! ")
	}
}
