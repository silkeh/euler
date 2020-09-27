package consumer

// Consumer consumes numbers in a certain way.
type Consumer interface {
	// Run against an input channel until it is closed..
	Run(in <-chan int, done chan<- bool)

	// Result returns the result of the consumption.
	Result() int
}
