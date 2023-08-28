package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Qual é o número? ")
	fmt.Scan(&n1)
	var n2 int
	n2 = n1 - 1
	fmt.Println("Seu antecessor é: ", n2)
	var n3 int
	n3 = n1 + 1
	fmt.Println("Seu sucessor é: ", n3)
}
