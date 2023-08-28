package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Quantos dias trabalhados? ")
	fmt.Scan(&n1)
	var n2 int
	fmt.Print("Qual é o valor da diária? ")
	fmt.Scan(&n2)
	var n3 int
	n3 = n1 * n2
	fmt.Println("Seu salário é de: ", n3)
}
