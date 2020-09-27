package filter

// Filter filters incoming numbers.
type Filter interface {
	// Run against an input and output channel until the input channel is closed.
	// The output channel is closed on return.
	Run(in <-chan int, out chan<- int)
}
