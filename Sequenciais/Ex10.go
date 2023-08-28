package main

import "fmt"

func main() {
	var kg float64
	fmt.Print("Qual é o seu peso em kg? ")
	fmt.Scan(&kg)
	var libra float64
	libra = kg * 2.2
	fmt.Println("O seu peso em libras é de: ", libra)
}
