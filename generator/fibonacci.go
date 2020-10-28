package generator

// Fibonacci generates the Fibonacci sequence.
type Fibonacci struct {
	f0, f1, n, max int
}

// NewFibonacci creates a Fibonacci Generator.
// The generates either `n` numbers, or until `max` is exceeded (whichever comes first).
func NewFibonacci(f0, f1, n, max int) Generator {
	return &Fibonacci{f0: f0, f1: f1, n: n, max: max}
}

// Run against a channel.
// The channel is closed when the Sequence is done.
func (f *Fibonacci) Run(out chan<- int, done <-chan bool) {
	defer close(out)

	out <- f.f0
	for i := 0; i < f.n; i++ {
		select {
		case <-done:
			return
		default:
			out <- f.f1

			next := f.f0 + f.f1
			if next > f.max {
				return
			}

			f.f0, f.f1 = f.f1, next
		}
	}
}
