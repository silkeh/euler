package filter

import (
	bigFilter "github.com/silkeh/euler/filter/big"
	"math/big"
	"sync"
)

// BigInt defines a Filter that can handle big.Int.
type BigInt struct {
	filters []bigFilter.Filter
}

// NewBigInt returns a Filter that internally works with *big.Int, allowing values bigger than 2^64.
// Automatically parses int to *big.Int and vice versa.
// bigFilters defines the list of BigFilters where the *big.Int is parsed through
func NewBigInt(bigFilters ...bigFilter.Filter) Filter {
	return &BigInt{
		filters: bigFilters,
	}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (bi *BigInt) Run(in <-chan int, out chan<- int) {
	var wg sync.WaitGroup
	wg.Add(len(bi.filters) + 2)

	channels := make([]chan *big.Int, len(bi.filters)+1)
	channels[0] = make(chan *big.Int)

	// Run conversion from int to *big.Int in goroutine
	go func() {
		ConvertTo(in, channels[0])
		wg.Done()
	}()

	for i := range bi.filters {
		channels[i+1] = make(chan *big.Int)

		// Copy i to local scope to avoid it changing before the filter starts
		j := i

		// Run the filter in a goroutine
		go func() {
			bi.filters[j].Run(channels[j], channels[j+1])
			wg.Done()
		}()
	}

	// Run conversion from *big.Int to int in goroutine
	go func() {
		ConvertFrom(channels[len(bi.filters)], out)
		wg.Done()
	}()

	wg.Wait()
}

// ConvertTo *big.Int from int until channel is closed.
// The output channel is closed on return.
func ConvertTo(in <-chan int, out chan<- *big.Int) {
	defer close(out)
	for v := range in {
		out <- big.NewInt(int64(v))
	}
}

// ConvertFrom big.Int to int until channel is closed.
// The output channel is closed on return.
func ConvertFrom(in <-chan *big.Int, out chan<- int) {
	defer close(out)
	for v := range in {
		out <- int(v.Int64())
	}
}
