package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Qual é o preço do produto? ")
	fmt.Scan(&n1)
	var n2 int
	n2 = n1 * 10 / 100
	var n3 int
	n3 = n1 - n2
	fmt.Println("O preço com o desconto é de: ", n3)
}
