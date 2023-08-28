package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Qual é a sua idade? ")
	fmt.Scan(&n1)
	var n2 int
	n2 = n1 * 365
	fmt.Print("Sua idade em dias é: ", n2)
}
