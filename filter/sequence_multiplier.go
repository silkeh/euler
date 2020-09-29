package filter

// SequenceMultiplier filters palindrome numbers.
type SequenceMultiplier struct {
	history []int
}

// NewSequenceMultiplier creates a palindrome-finding Filter.
func NewSequenceMultiplier(n int) Filter {
	history := make([]int, n)
	return &SequenceMultiplier{history: history}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (f *SequenceMultiplier) Run(in <-chan int, out chan<- int) {
	defer close(out)

	for v := range in {
		product := v
		for i := 1; i < len(f.history); i++ {
			product *= f.history[i]
			f.history[i-1] = f.history[i]
		}
		out <- product
		f.history[len(f.history)-1] = v
	}
}
