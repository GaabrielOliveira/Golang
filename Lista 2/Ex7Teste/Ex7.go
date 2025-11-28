package main

import (
	"fmt"
	"sort"
)

func OrdenaString(texto string) string {
	chars := []rune(texto)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func main() {
	var texto string
	fmt.Print("Digite uma sequência de caracteres: ")
	fmt.Scan(&texto)

	textoOrdenado := OrdenaString(texto)

	fmt.Println("Caracteres ordenados:", textoOrdenado)
}