// Arquivo: Ex2_test.go
package main

import "testing"

func BenchmarkDivisao(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divisao(100, 5)
	}
}