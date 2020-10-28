package condition

// equal compares the received value.
// Returns True if received is greater than or equal to threshold.
type equal struct {
	threshold int
}

// NewEqual returns a greater-then or equal to condition pipeline.
func NewEqual(threshold int) Condition {
	return &equal{threshold: threshold}
}

// Run against an input channel until it is closed.
func (ge *equal) Run(in <-chan int, out chan<- bool) {
	defer close(out)

	for v := range in {
		out <- v == ge.threshold
	}
}
