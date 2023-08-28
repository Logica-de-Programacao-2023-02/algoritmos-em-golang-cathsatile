package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Qual é o número? ")
	fmt.Scan(&n1)
	var n2 int
	n2 = n1 * 2
	fmt.Println("O dobro do número é: ", n2)
	var n3 int
	n3 = n1 * 3
	fmt.Println("O triplo do número é: ", n3)
	var n4 int
	n4 = n1 * 4
	fmt.Println("O quadruplo do número é: ", n4)
}
