package consumer

// Summer sums the received values.
// It is not thread-safe.
type Summer struct {
	sum int
}

// NewSummer returns a summing Consumer.
func NewSummer() Consumer {
	return &Summer{}
}

// Run against an input channel until it is closed..
func (s *Summer) Run(in <-chan int, done chan<- bool) {
	for v := range in {
		s.sum += v
	}
}

// Result returns the result of the consumption.
func (s *Summer) Result() int {
	return s.sum
}
