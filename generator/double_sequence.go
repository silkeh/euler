package generator

// doubleSequence generates a sequence of numbers.
type doubleSequence struct {
	min, max, inc int
}

// NewDoubleSequence creates a doubleSequence Generator.
func NewDoubleSequence(min, max, inc int) Generator {
	return &doubleSequence{min: min, max: max, inc: inc}
}

// Run against a channel.
// The channel is closed when the doubleSequence is done.
func (s *doubleSequence) Run(out chan<- int, done <-chan bool) {
	defer close(out)

	for i := s.min; s.checkEnd(i); i += s.inc {
		for j := s.min; s.checkEnd(j); j += s.inc {
			select {
			case <-done:
				return
			default:
				out <- i
				out <- j
			}
		}
	}
}

// checkEnd returns true if the i has reached max.
func (s *doubleSequence) checkEnd(i int) bool {
	if s.inc < 0 {
		return i >= s.max
	}
	return i <= s.max
}
