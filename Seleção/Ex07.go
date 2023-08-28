package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Qual é o seu salário? ")
	fmt.Scan(&n1)
	var n2 int
	n2 = n1 * 110 / 100
	var n3 int
	n3 = n1 * 105 / 100
	if n1 < 1000 {
		fmt.Print("Seu novo salário é: ", n2)
	} else {
		fmt.Println("Seu novo salário é: ", n3)
	}
}
