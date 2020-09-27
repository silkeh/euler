package consumer

import "math"

// Maxer finds the maximum received value.
// It is not thread-safe.
type Maxer struct {
	max int
}

// NewMaxer returns a max-finding Consumer.
func NewMaxer() Consumer {
	return &Maxer{max: math.MinInt64}
}

// Run against an input channel until it is closed..
func (m *Maxer) Run(in <-chan int, done chan<- bool) {
	for v := range in {
		if v > m.max {
			m.max = v
		}
	}
}

// Result returns the result of the consumption.
func (m *Maxer) Result() int {
	return m.max
}
