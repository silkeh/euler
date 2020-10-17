package util

// split one channel in to multiple output channels.
// Each output is a direct copy of the input channel
type split struct{}

// NewSplit returns a channel split
func NewSplit() Multiplex {
	return &split{}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (s split) Run(in <-chan int, out ...chan<- int) {
	defer func() {
		// close all output channels
		for _, v := range out {
			close(v)
		}
	}()

	for vi := range in {
		for _, vo := range out {
			vo <- vi
		}
	}
}
