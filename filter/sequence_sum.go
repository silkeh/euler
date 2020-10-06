package filter

// SequenceSum is a filter with a sequential sum
type SequenceSum struct {
	n int
}

// NewSequenceSum creates sequential sum Filter.
func NewSequenceSum(n int) Filter {
	return &SequenceSum{n: n}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (s *SequenceSum) Run(in <-chan int, out chan<- int) {
	defer close(out)
	for {
		sum := 0
		for i := 0; i < s.n; i++ {
			v, ok := <-in
			if !ok {
				return
			}
			sum += v
		}
		out <- sum
	}
}
