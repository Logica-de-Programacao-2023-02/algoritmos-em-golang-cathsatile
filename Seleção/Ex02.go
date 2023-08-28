package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Print("Qual é o primeiro número? ")
	fmt.Scan(&n1)
	fmt.Print("Qual é o segundo número? ")
	fmt.Scan(&n2)
	fmt.Print("Qual é o terceiro número? ")
	fmt.Scan(&n3)
	if n1 < n2 && n1 < n3 {
		fmt.Println("O menor número é: ", n1)
	}
	if n2 < n1 && n2 < n3 {
		fmt.Println("O menor número é: ", n2)
	}
	if n3 < n1 && n3 < n2 {
		fmt.Println("O menor número é: ", n3)
	}
}
