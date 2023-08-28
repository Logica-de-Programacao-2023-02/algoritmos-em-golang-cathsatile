package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Qual é o primeiro número? ")
	fmt.Scan(&n1)
	var n2 int
	fmt.Print("Qual é o segunda número? ")
	fmt.Scan(&n2)
	var n3 int
	fmt.Print("Qual é o terceiro número? ")
	fmt.Scan(&n3)
	var n4, n5, n6 int
	n4 = n1 * 2
	n5 = n2 * 3
	n6 = n3 * 5
	var n7 int
	n7 = n4 + n5 + n6
	var n8 int
	n8 = n7 / 10
	fmt.Println("A média ponderada é: ", n8)
}
