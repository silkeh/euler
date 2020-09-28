package generator

import "git.slxh.eu/go/euler/filter"

// Prime generates a sequence of numbers.
type Prime struct {
	n, max      int
	knownPrimes []int
}

// NewPrime creates a Prime Generator.
func NewPrime(n, max int) Generator {
	p := make([]int, 1, 10000)
	p[0] = 2
	return &Prime{n: n, max: max, knownPrimes: p}
}

// Run against a channel.
// The channel is closed when the Prime is done.
func (s *Prime) Run(out chan<- int, done <-chan bool) {
	defer close(out)

	n := 1
	out <- 2
	for i := 3; n <= s.n && i <= s.max; i += 2 {
		select {
		case <-done:
			return
		default:
			if filter.IsPrime(i) {
				n++
				out <- i
			}
		}
	}
}
