package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Qual é o número? ")
	fmt.Scan(&n1)
	if n1%3 == 0 && n1%5 == 0 {
		fmt.Println("O número é múltiplo de 3 e de 5! ")
	} else {
		fmt.Println("O número NÃO é múltiplo de 3 e de 5! ")
	}
}
