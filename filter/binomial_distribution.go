package filter

import (
	"math/big"
)

// binomialDistribution calculates the BDF of 2 given values.
type binomialDistribution struct{}

// NewBinomialDistribution creates a binomial distribution filter.
func NewBinomialDistribution() Filter {
	return &binomialDistribution{}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (b *binomialDistribution) Run (in <-chan int, out chan<- int){
	defer close(out)
	for {
		n, ok1 := <-in
		k, ok2 := <-in
		if !(ok1 && ok2) {
			return
		}
		out <- bdf(n,k)
	}
}

// bdf calculates the binomial distribution function on n, k.
func bdf (n, k int) int {
	out := factorial(n)
	out.Div(out,factorial(n-k))
	out.Div(out,factorial(k))
	return int(out.Int64())
}

// factorial calculates the factorial of in.
// returns a big.Int pointer.
func factorial(in int) (out *big.Int) {
	out = new(big.Int)
	out.SetInt64(1)
	j := new(big.Int)
	for i := 2; i <= in; i++ {
		j.SetInt64(int64(i))
		out.Mul(out,j)
	}
	return
}