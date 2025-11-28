package main

import "fmt"

func Soma(a, b int) int {
	return a + b
}

func main() {
	var a, b int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&a)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&b)

	soma := Soma(a, b)

	fmt.Println("O resultado da soma é: ", soma)
}