package condition

// less compares the received value.
// Returns True if received is less than threshold.
type less struct {
	threshold int
}

// NewLess returns a less-then condition pipeline.
func NewLess(threshold int) Condition {
	return &less{threshold: threshold}
}

// Run against an input channel until it is closed.
func (l *less) Run(in <-chan int, out chan<- bool) {
	defer close(out)

	for v := range in {
		out <- v < l.threshold
	}
}
