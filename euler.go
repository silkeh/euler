package main

import (
	"fmt"
	"git.slxh.eu/go/euler/consumer"
	"git.slxh.eu/go/euler/filter"
	"git.slxh.eu/go/euler/generator"
	"math"
	"time"
)

var problems = []func() int{
	// Problem 1
	func() int {
		return RunPipeline(
			generator.NewSequence(1, 999, 1),
			consumer.NewSummer(),
			filter.NewThreeOrFiveFilter(),
		)
	},

	// Problem 2
	func() int {
		return RunPipeline(
			generator.NewFibonacci(math.MaxInt64, 4000000),
			consumer.NewSummer(),
			filter.NewEven(),
		)
	},

	// Problem 3
	func() int {
		v := 600851475143
		return RunPipeline(
			generator.NewPrime(math.MaxInt64, v),
			consumer.NewFactorizer(v),
		)
	},

	// Problem 4
	func() int {
		return RunPipeline(
			generator.NewProducts(900, 900, 999, 999),
			consumer.NewMaxer(false),
			filter.NewPalindrome(),
		)
	},

	// Problem 5
	func() int {
		// the step size is the product of all primes
		step := 1*2*3*5*7*11*13*17*19
		return RunPipeline(
			generator.NewSequence(step, math.MaxInt64, step),
			consumer.NewMiner(true),
			// the filters are the highest unique prime multiples
			filter.NewFractor(12),
			filter.NewFractor(14),
			filter.NewFractor(15),
			filter.NewFractor(16),
			filter.NewFractor(18),
			filter.NewFractor(20),
		)
	},
}

func main() {
	for i, p := range problems {
		start := time.Now()
		fmt.Printf("Problem %03d: %v (%s)\n", i+1, p(), time.Since(start))
	}
}
