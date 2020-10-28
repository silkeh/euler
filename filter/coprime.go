package filter

// Coprime filters Coprime numbers
type Coprime struct{}

// NewCoprime creates a Coprime-finding Filter.
func NewCoprime() Filter {
	return &Coprime{}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (p *Coprime) Run(in <-chan int, out chan<- int) {
	defer close(out)
	for {
		v1, ok1 := <-in
		v2, ok2 := <-in
		if !(ok1 && ok2) {
			return
		}
		if isCoprime(v1, v2) {
			out <- v1
			out <- v2
		}
	}
}

// isCoprime returns true when a and b are coprime.
func isCoprime(a, b int) bool {
	min := b
	if a < b {
		min = a
	}
	for i := 2; i < min; i++ {
		if a%i == 0 && b%i == 0 {
			return false
		}
	}
	return true
}
