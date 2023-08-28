package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Qual é o primeiro número? ")
	fmt.Scan(&n1)
	fmt.Print("Qual é o segundo número? ")
	fmt.Scan(&n2)
	if n1 > 0 && n2 > 0 {
		fmt.Println("A multiplicação dos números é: ", n1*n2)
	}
	if n1 < 0 && n2 > 0 || n1 > 0 && n2 < 0 {
		fmt.Println("A soma dos números é: ", n1+n2)
	}
}
