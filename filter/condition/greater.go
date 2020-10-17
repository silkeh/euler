package condition

// greater compares the received value.
// Returns True if received is greater than threshold.
type greater struct {
	threshold int
}

// NewGreater returns a greater-than condition pipeline.
func NewGreater(threshold int) Condition {
	return &greater{threshold: threshold}
}

// Run against an input channel until it is closed.
func (g *greater) Run(in <-chan int, out chan<- bool) {
	defer close(out)

	for v := range in {
		out <- v > g.threshold
	}
}
