// Arquivo: Ex6_test.go
package main

import (
	"math/rand"
	"testing"
)

func BenchmarkOrdenaNumeros(b *testing.B) {
	numerosParaOrdenar := make([]int, 1000)
	for i := 0; i < len(numerosParaOrdenar); i++ {
		numerosParaOrdenar[i] = rand.Intn(10000)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
        b.StopTimer()
		copia := make([]int, len(numerosParaOrdenar))
		copy(copia, numerosParaOrdenar)
		b.StartTimer()

		OrdenaNumeros(copia)
	}
}