package condition

// Condition compares received values with given statement.
type Condition interface {
	// Run against an input and output channel until the input channel is closed.
	// The output channel is closed on return.
	Run(in <-chan int, out chan<- bool)
}
