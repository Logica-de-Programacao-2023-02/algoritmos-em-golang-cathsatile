package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Qual é o número? ")
	fmt.Scan(&n1)
	if n1%2 == 0 {
		fmt.Println("O número é: Par")
	} else {
		fmt.Println("O número é: Ímpar")
	}
}
