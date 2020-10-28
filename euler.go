package main

import (
	"fmt"
	"github.com/silkeh/euler/consumer"
	"github.com/silkeh/euler/filter"
	"github.com/silkeh/euler/filter/condition"
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
			generator.NewFibonacci(0, 1, math.MaxInt64, 4000000),
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
		digits := "73167176531330624919225119674426574742355349194934" +
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
			"71636269561882670428252483600823257530420752963450"
		return RunPipeline(
			generator.NewDigitReader(digits, 1),
			consumer.NewMaxer(false),
			filter.NewSequenceMultiplier(13),
		)
	},

	// Problem 9
	func() int {
		return RunPipeline(
			generator.NewDoubleSequence(1, 100, 1),
			consumer.NewMaxer(true),
			filter.NewCoprime(),
			filter.NewPythagoreanTripleter(25),
			filter.NewConditional(
				filter.NewProducter(3),
				filter.NewSequenceSum(3),
				condition.NewEqual(1000),
				1,
			),
		)
	},

	// Problem 10
	func() int {
		return RunPipeline(
			generator.NewPrime(math.MaxInt64, 2000000),
			consumer.NewSummer())
	},

	// Problem 11
	func() int {
		grid := []string{
			"0802229738150040007504050778521250779108",
			"4949994017811857608717409843694804566200",
			"8149317355791429937140675388300349133665",
			"5270952304601142692468560132567137023691",
			"2231167151676389419236542240402866331380",
			"2447326099034502447533537836842035171250",
			"3298812864236710263840675954706618386470",
			"6726206802621220956394396308409166499421",
			"2455580566739926971778789683148834896372",
			"2136230975007644204535140061339734313395",
			"7817532822753167159403800462161409535692",
			"1639054296353147555888240017542436298557",
			"8656004835718907054444374460215851541758",
			"1980816805944769287392138652177704895540",
			"0452088397359916079757321626267933279866",
			"8836688757622072034633674655123263935369",
			"0442167338253911249472180846293240627636",
			"2069364172302388346299698267598574043616",
			"2073352978319001743149714886811623570554",
			"0170547183515469169233486143520189196748",
		}
		return RunPipeline(
			generator.NewDirectionGenerator(20, 20, 4),
			consumer.NewMaxer(false),
			filter.NewSelectDigit(grid, 2),
			filter.NewProducter(4),
		)
	},

	// Problem 12
	func() int {
		return RunPipeline(
			generator.NewSequence(1, 100000, 1),
			consumer.NewMaxer(true),
			filter.NewCumSummer(),
			filter.NewConditional(
				filter.NewScale(1),
				filter.NewDivisorCount(),
				condition.NewGreaterEqual(500),
				1,
			),
		)
	},
}

func main() {
	for i, p := range problems {
		start := time.Now()
		fmt.Printf("Problem %03d: %v (%s)\n", i+1, p(), time.Since(start))
	}
}
