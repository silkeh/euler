package filter

// Producter multiplies a set of n numbers.
type Producter struct {
	n int
}

// NewProducter creates a n multiply Filter.
func NewProducter(n int) Filter {
	return &Producter{n: n}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (p *Producter) Run(in <-chan int, out chan<- int) {
	defer close(out)
	i := 0
	product := 1
	for v := range in {
		product *= v
		i++
		if !(i < p.n) {
			out <- product
			product = 1
			i = 0
		}
	}
}
