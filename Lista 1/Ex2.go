// Arquivo: Ex2.go
package main

import "fmt"

func Divisao(a, b int) int {
	return a / b
}

func main() {
	var a, b int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&a)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&b)

	div := Divisao(a, b)

	fmt.Println("O resultado da divisão é: ", div)
}