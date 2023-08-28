package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Qual é o primeiro número? ")
	fmt.Scan(&n1)
	var n2 int
	fmt.Print("Qual é o segundo número? ")
	fmt.Scan(&n2)
	var n3 int
	fmt.Print("Qual é o terceiro número? ")
	fmt.Scan(&n3)
	var n4 int
	n4 = n1 + n2 + n3
	fmt.Print("A soma dos números é: ", n4)
}
