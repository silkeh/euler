package main

import (
	"testing"
)

var answers = []int{
	// 1
	233168,
	4613732,
	6857,
	906609,
	232792560,
	25164150,
	104743,
	23514624000,
	0, // skipped
	// 10
	142913828922,
	70600674,
}

func Test(t *testing.T) {
	if len(problems) > len(answers) {
		t.Errorf("Missing %v answer(s)", len(problems)-len(answers))
	}

	for i, answer := range answers {
		result := problems[i]()
		if answer != result {
			t.Errorf("Incorrect solution for problem %03d: expected %v, got %v", i+1, answer, result)
		}
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
