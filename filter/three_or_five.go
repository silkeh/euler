package filter

// ThreeOrFiveFilter filters numbers that are a multiple of three or five.
type ThreeOrFiveFilter struct{}

// NewThreeOrFiveFilter creates a three or five Filter.
func NewThreeOrFiveFilter() Filter {
	return &ThreeOrFiveFilter{}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (f *ThreeOrFiveFilter) Run(in <-chan int, out chan<- int) {
	defer close(out)

	for v := range in {
		switch {
		case v%3 == 0:
			out <- v
		case v%5 == 0:
			out <- v
		default:
			// do nothing
		}
	}
}
