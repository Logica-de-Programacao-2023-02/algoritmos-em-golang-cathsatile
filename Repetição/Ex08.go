package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Qual é o número?")
	fmt.Scan(&n1)
	for i := n1; i > 0; i-- {
		if n1%i == 0 {
			fmt.Println(i)
		}
	}

}
