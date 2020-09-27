package main

import (
	"fmt"
	"git.slxh.eu/go/euler/consumer"
	"git.slxh.eu/go/euler/filter"
	"git.slxh.eu/go/euler/generator"
	"math"
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
			generator.NewSequence(1, int(math.Sqrt(float64(v))), 1),
			consumer.NewMaxer(),
			filter.NewFactor(v),
			filter.NewPrime(),
		)
	},

	// Problem 4
	func() int {
		return RunPipeline(
			generator.NewProducts(100, 100, 999, 999),
			consumer.NewMaxer(),
			filter.NewPalindrome(),
		)
	},

	// Problem 5
	func() int {
		max := 20
		maxValue := 1
		maxPrime := 0

		// Create filters for all numbers from 2 to `max`.
		filters := make([]filter.Filter, 0, max)
		for i := 2; i <= max; i++ {
			filters = append(filters, filter.NewFractor(i))

			// Multiply to get the upper search bound.
			maxValue *= i

			// Get the highest prime number to use as step.
			if filter.IsPrime(i) {
				maxPrime = i
			}
		}

		return RunPipeline(
			generator.NewSequence(maxPrime, int(math.Sqrt(float64(maxValue))), maxPrime),
			consumer.NewMiner(true),
			filters...,
		)
	},
}

func main() {
	for i, p := range problems {
		fmt.Printf("Problem %03d: %v\n", i+1, p())
	}
}
