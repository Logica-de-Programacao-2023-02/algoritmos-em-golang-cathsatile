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
	if n1 <= n2 && n1 <= n3 && n2 <= n3 {
		fmt.Print("A ordem crescente dos números é: ", n1, n2, n3)
	}
	if n1 <= n2 && n1 <= n3 && n3 <= n2 {
		fmt.Print("A ordem crescente dos números é: ", n1, n3, n2)
	}
	if n2 <= n1 && n2 <= n3 && n1 <= n3 {
		fmt.Print("A ordem crescente dos números é: ", n2, n1, n3)
	}
	if n2 <= n1 && n2 <= n3 && n3 <= n1 {
		fmt.Print("A ordem crescente dos números é: ", n2, n3, n1)
	}
	if n3 <= n2 && n3 <= n1 && n1 <= n2 {
		fmt.Print("A ordem crescente dos números é: ", n3, n1, n2)
	}
	if n3 <= n1 && n3 <= n2 && n2 <= n1 {
		fmt.Print("A ordem crescente dos números é: ", n3, n2, n1)
	}

}
