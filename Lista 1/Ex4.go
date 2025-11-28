// Arquivo: Ex4.go
package main

import "fmt"

func VerificaNumero(num int) string {
	if num == 0 {
		return "O numero é zero (Nem positivo nem negativo)."
	} else {
		if num > 0 && num%2 == 0 {
			return "O número é positivo e par"
		} else if num > 0 && num%2 != 0 {
			return "O número é positivo e impar"
		} else if num < 0 && num%2 == 0 {
			return "O número é negativo e par"
		} else {
			return "O número é negativo e impar"
		}
	}
}

func main() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	resultado := VerificaNumero(num)
	fmt.Println(resultado)
}