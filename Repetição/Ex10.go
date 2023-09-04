package main

import "fmt"

func main() {
	var x, maior int
	x = -1
	for x != 0 {
		fmt.Print("Digite um número: ")
		fmt.Scan(&x)

		if x > maior {
			maior = x
		}

	}
	fmt.Println("O maior número é: ", maior)

}
