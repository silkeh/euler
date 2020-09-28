package consumer

// Factorizer sums the received values.
// It is not thread-safe.
type Factorizer struct {
	n int
	last int
}

// NewFactorizer returns a factorizing Consumer.
func NewFactorizer(n int) Consumer {
	return &Factorizer{n: n}
}

// Run against an input channel until it is closed..
func (s *Factorizer) Run(in <-chan int, done chan<- bool) {
	for v := range in {
		for s.n % v == 0 {
			s.n /= v
			s.last = v
		}

		if s.n <= 1 {
			done <- true
		}
	}
}

// Result returns the result of the consumption.
func (s *Factorizer) Result() int {
	return s.last
}
