package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Qual é o primeiro número? ")
	fmt.Scan(&n1)
	fmt.Print("Qual é o segundo número? ")
	fmt.Scan(&n2)
	if n1 > n2 {
		fmt.Println("O maior número é: ", n1)
	} else {
		fmt.Println("O maior número é: ", n2)
	}

}
