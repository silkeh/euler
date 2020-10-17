package util

// Multiplex one input into multiple output channels
type Multiplex interface {
	Run(in <-chan int, out ...chan<- int)
}
