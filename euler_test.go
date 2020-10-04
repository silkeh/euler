package main

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	for i, p := range problems {
		start := time.Now()
		t.Logf("Problem %03d: %v (%s)\n", i+1, p(), time.Since(start))
	}
}

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
