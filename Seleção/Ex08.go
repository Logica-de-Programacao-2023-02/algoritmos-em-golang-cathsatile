package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Digite um número de 1 a 7: ")
	fmt.Scan(&n1)
	switch n1 {
	case 1:

		fmt.Println("O dia da semana é Domingo")

	case 2:

		fmt.Println("O dia da semana é Segunda-feira")

	case 3:

		fmt.Println("O dia da semana é Terça-feira")

	case 4:

		fmt.Println("O dia da semana é Quarta-feira")

	case 5:

		fmt.Println("O dia da semana é Quinta-feira")

	case 6:

		fmt.Println("O dia da semana é Sexta-feira")

	case 7:

		fmt.Println("O dia da semana é Domingo")

	}
}
