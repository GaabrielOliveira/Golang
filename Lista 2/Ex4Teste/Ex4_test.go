// Arquivo: Ex4_test.go
package main

import "testing"

func BenchmarkVerificaNumero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VerificaNumero(-42) // Testamos com um número negativo e par
	}
}