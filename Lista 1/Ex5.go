package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	if isPrime(num) {
		fmt.Printf("%d é primo\n", num)
	} else {
		fmt.Printf("%d não é primo\n", num)
	}
}
