package main

import (
	"fmt"
	"github.com/silkeh/euler/consumer"
	"github.com/silkeh/euler/filter"
	"github.com/silkeh/euler/generator"
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
		step := 1 * 2 * 3 * 5 * 7 * 11 * 13 * 17 * 19
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

	// Problem 6
	func() int {
		n := 100
		return int(math.Pow(float64(RunPipeline(
			generator.NewSequence(1, n, 1),
			consumer.NewSummer(),
		)), 2)) - RunPipeline(
			generator.NewSequence(1, n, 1),
			consumer.NewSummer(),
			filter.NewPowerer(2),
		)
	},

	// Problem 7
	func() int {
		return RunPipeline(
			generator.NewPrime(10001, math.MaxInt64),
			consumer.NewMaxer(false),
		)
	},

	// Problem 8
	func() int {
		digits := []byte("73167176531330624919225119674426574742355349194934" +
			"96983520312774506326239578318016984801869478851843" +
			"85861560789112949495459501737958331952853208805511" +
			"12540698747158523863050715693290963295227443043557" +
			"66896648950445244523161731856403098711121722383113" +
			"62229893423380308135336276614282806444486645238749" +
			"30358907296290491560440772390713810515859307960866" +
			"70172427121883998797908792274921901699720888093776" +
			"65727333001053367881220235421809751254540594752243" +
			"52584907711670556013604839586446706324415722155397" +
			"53697817977846174064955149290862569321978468622482" +
			"83972241375657056057490261407972968652414535100474" +
			"82166370484403199890008895243450658541227588666881" +
			"16427171479924442928230863465674813919123162824586" +
			"17866458359124566529476545682848912883142607690042" +
			"24219022671055626321111109370544217506941658960408" +
			"07198403850962455444362981230987879927244284909188" +
			"84580156166097919133875499200524063689912560717606" +
			"05886116467109405077541002256983155200055935729725" +
			"71636269561882670428252483600823257530420752963450")
		return RunPipeline(
			generator.NewDigitReader(digits),
			consumer.NewMaxer(false),
			filter.NewSequenceMultiplier(13),
		)
	},

	// Problem 9
	func() int {
		return 0
	},

	// Problem 10
	func() int {
		return RunPipeline(
			generator.NewPrime(math.MaxInt64, 2000000),
			consumer.NewSummer())
	},
}

func main() {
	for i, p := range problems {
		start := time.Now()
		fmt.Printf("Problem %03d: %v (%s)\n", i+1, p(), time.Since(start))
	}
}
