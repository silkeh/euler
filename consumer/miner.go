package consumer

import (
	"math"
)

// Miner finds the maximum received value.
// It is not thread-safe.
type Miner struct {
	min       int
	stopFirst bool
}

// NewMiner returns a min-finding Consumer.
func NewMiner(stopOnFirst bool) Consumer {
	return &Miner{min: math.MaxInt64, stopFirst: stopOnFirst}
}

// Run against an input channel until it is closed..
func (m *Miner) Run(in <-chan int, done chan<- bool) {
	for v := range in {
		if v < m.min {
			m.min = v
			if m.stopFirst {
				done <- true
			}
		}
	}
}

// Result returns the result of the consumption.
func (m *Miner) Result() int {
	return m.min
}
