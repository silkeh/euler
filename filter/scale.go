package filter

// scale the received values
type scale struct {
	s int
}

// NewScale returns a scaling filter
func NewScale(s int) Filter {
	return &scale{s: s}
}

// Run against an input channel until it is closed.
func (s *scale) Run(in <-chan int, out chan<- int) {
	defer close(out)

	for v := range in {
		out <- v * s.s
	}
}
