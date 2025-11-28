package main

import "fmt"

func main() {
	var n int

	fmt.Print("Digite um número: ")
	fmt.Scanln(&n)

	antecessor := n - 1
	sucessor := n + 1

	fmt.Println(n)
	fmt.Printf("Antecessor: %d\n", antecessor)
	fmt.Printf("Sucessor: %d\n", sucessor)
}
