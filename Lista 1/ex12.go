package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&a)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&b)

	igualdade := a == b

	fmt.Print("Os vaores são iguais?\n", igualdade)
}