package condition

// greaterEqual compares the received value.
// Returns True if received is greater than or equal to threshold.
type greaterEqual struct {
	threshold int
}

// NewGreaterEqual returns a greater-than or equal to condition pipeline.
func NewGreaterEqual(threshold int) Condition {
	return &greaterEqual{threshold: threshold}
}

// Run against an input channel until it is closed.
func (ge *greaterEqual) Run(in <-chan int, out chan<- bool) {
	defer close(out)

	for v := range in {
		out <- v >= ge.threshold
	}
}
