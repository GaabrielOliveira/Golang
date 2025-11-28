// Arquivo: Ex6.go
package main

import (
	"fmt"
	"sort"
)

func OrdenaNumeros(numeros []int) {
	sort.Ints(numeros)
}

func main() {
	var n int
	fmt.Print("Digite a quantidade de números: ")
	fmt.Scan(&n)

	numeros := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&numeros[i])
	}

	OrdenaNumeros(numeros)

	fmt.Println("Sequência ordenada:", numeros)
}