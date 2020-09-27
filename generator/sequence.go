package generator

// Sequence generates a sequence of numbers.
type Sequence struct {
	min, max, inc int
}

// NewSequence creates a Sequence Generator.
func NewSequence(min, max, inc int) Generator {
	return &Sequence{min: min, max: max, inc: inc}
}

// Run against a channel.
// The channel is closed when the Sequence is done.
func (s *Sequence) Run(out chan<- int, done <-chan bool) {
	defer close(out)
	for i := s.min; i <= s.max; i += s.inc {
		select {
		case <-done:
			return
		default:
			out <- i
		}
	}
}
