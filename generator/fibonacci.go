package generator

// Fibonacci generates the Fibonacci sequence.
type Fibonacci struct {
	n, max int
}

// NewFibonacci creates a Fibonacci Generator.
// The generates either `n` numbers, or until `max` is exceeded (whichever comes first).
func NewFibonacci(n, max int) Generator {
	return &Fibonacci{n: n, max: max}
}

// Run against a channel.
// The channel is closed when the Sequence is done.
func (f *Fibonacci) Run(out chan<- int, done <-chan bool) {
	defer close(out)

	prev0 := 0
	prev1 := 1
	for i := 0; i < f.n; i++ {
		select {
		case <-done:
			return
		default:
			out <- prev1

			next := prev0 + prev1
			if next > f.max {
				return
			}

			prev0, prev1 = prev1, next
		}
	}
}
