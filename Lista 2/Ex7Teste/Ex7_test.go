// Arquivo: Ex7_test.go
package main

import "testing"

func BenchmarkOrdenaString(b *testing.B) {
	texto := "umasequenciadecaracteresparaserordenada"
	b.ResetTimer() // Garante que a atribuição da string não seja medida

	for i := 0; i < b.N; i++ {
		OrdenaString(texto)
	}
}