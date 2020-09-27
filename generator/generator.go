package generator

// Generator is a pipeline component that generates numbers.
type Generator interface {
	// Run against a channel.
	// The channel is closed when the Sequence is done.
	Run(out chan<- int, done <-chan bool)
}
