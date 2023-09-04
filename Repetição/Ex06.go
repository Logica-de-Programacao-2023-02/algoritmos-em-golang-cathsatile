package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Qual é o número?")
	fmt.Scan(&n1)
	for i := 1; i <= 10; i++ {
		fmt.Println(i * n1)
	}

}
