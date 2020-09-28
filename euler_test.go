package main

import "testing"

func benchmarkEulerN(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		problems[n]()
	}
}

func BenchmarkEuler1(b *testing.B) {
	benchmarkEulerN(0, b)
}

func BenchmarkEuler2(b *testing.B) {
	benchmarkEulerN(1, b)
}

func BenchmarkEuler3(b *testing.B) {
	benchmarkEulerN(2, b)
}

func BenchmarkEuler4(b *testing.B) {
	benchmarkEulerN(3, b)
}
