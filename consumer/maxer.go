package consumer

import (
	"math"
)

// maxer finds the maximum received value.
// It is not thread-safe.
type maxer struct {
	max       int
	stopFirst bool
}

// NewMaxer returns a max-finding Consumer.
func NewMaxer(stopOnFirst bool) Consumer {
	return &maxer{max: math.MinInt64, stopFirst: stopOnFirst}
}

// Run against an input channel until it is closed..
func (m *maxer) Run(in <-chan int, done chan<- bool) {
	defer channelCleanup(in, done)
	for v := range in {
		if v > m.max {
			m.max = v
			if m.stopFirst {
				return
			}
		}
	}
}

// Result returns the result of the consumption.
func (m *maxer) Result() int {
	return m.max
}
