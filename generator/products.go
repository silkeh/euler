package generator

// Products generates a sequence of numbers that are products of each other.
type Products struct {
	min1, min2, max1, max2 int
}

// NewProducts creates a Products Generator.
func NewProducts(min1, min2, max1, max2 int) Generator {
	return &Products{min1: min1, min2: min2, max1: max1, max2: max2}
}

// Run against a channel.
// The channel is closed when the Products is done.
func (s *Products) Run(out chan<- int, done <-chan bool) {
	defer close(out)
	for i := s.min1; i <= s.max2; i++ {
		for j := s.min2; j <= s.max2; j++ {
			select {
			case <-done:
				return
			default:
				out <- i * j
			}
		}
	}
}
