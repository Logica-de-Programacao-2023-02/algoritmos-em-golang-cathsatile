package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Qual é o seu salário? ")
	fmt.Scan(&n1)
	var n2 int
	n2 = n1 * 115 / 100
	fmt.Print("Seu novo salário é: ", n2)
}
