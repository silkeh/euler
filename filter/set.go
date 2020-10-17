package filter

import (
	"sync"
)

// set is a Filter that contains 2 or more Filters
type set struct {
	filters []Filter
}

// NewSet returns a set of Filters
func NewSet(filters ...Filter) Filter {
	return &set{filters: filters}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (s *set) Run(in <-chan int, out chan<- int) {
	if len(s.filters) == 0 {
		// if there are no filters, create a no operations filter to connect in and out.
		s.filters = append(s.filters, NewNOP())
	}

	var wg sync.WaitGroup
	wg.Add(len(s.filters))
	// Create a list for all channels
	channelsIn := make([]<-chan int, len(s.filters))
	channelsOut := make([]chan<- int, len(s.filters))
	channelsIn[0] = in
	channelsOut[len(s.filters)-1] = out

	// Create all filters
	for i := range s.filters {
		// Create the output channel
		channel := make(chan int)
		if i < len(s.filters)-1 {
			channelsIn[i+1] = channel
			channelsOut[i] = channel
		}

		// Copy i to local scope to avoid it changing before the filter starts
		j := i

		// Run the filter in a goroutine
		go func() {
			s.filters[j].Run(channelsIn[j], channelsOut[j])
			wg.Done()
		}()
	}
	wg.Wait()
}
