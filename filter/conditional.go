package filter

import (
	"github.com/silkeh/euler/filter/condition"
	"github.com/silkeh/euler/util"
	"sync"
)

// conditional represents two parallel Filters and a condition.
// If the condition is triggered on the conditionFilter:
// Then n values from the dataFilter will be forwarded on the out channel.
// Else n values will be discarded.
type conditional struct {
	dataFilter      Filter
	conditionFilter Filter
	condition       condition.Condition
	n               int
}

// NewConditional creates a new conditional Filter
func NewConditional(dataFilter Filter, conditionFilter Filter, condition condition.Condition, n int) Filter {
	return &conditional{
		dataFilter:      dataFilter,
		conditionFilter: conditionFilter,
		condition:       condition,
		n:               n,
	}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (c *conditional) Run(in <-chan int, out chan<- int) {
	var wg sync.WaitGroup
	wg.Add(4)

	defer func() {
		close(out)
	}()

	split := util.NewSplit()
	channel1 := make(chan int)
	channel2 := make(chan int)
	dataChannel := make(chan int)
	conditionalChannel := make(chan int)
	booleanChannel := make(chan bool)

	// split in -> channel1; channel2
	go func() {
		split.Run(in, channel1, channel2)
		wg.Done()
	}()

	// run filter on data path
	go func() {
		c.dataFilter.Run(channel1, dataChannel)
		wg.Done()
	}()

	// run filter on conditional path
	go func() {
		c.conditionFilter.Run(channel2, conditionalChannel)
		wg.Done()
	}()

	// run condition on filtered conditional path
	go func() {
		c.condition.Run(conditionalChannel, booleanChannel)
		wg.Done()
	}()

	v := make([]int, c.n)
	ok := true
	// read from conditional channel
	for b := range booleanChannel {
		// read n values from input
		for i := 0; i < c.n; i++ {
			if v[i], ok = <-dataChannel; !ok {
				return
			}
		}
		if b {
			// push n values on output channel
			for i := 0; i < c.n; i++ {
				out <- v[i]
			}
		}
	}

	wg.Wait()
}
