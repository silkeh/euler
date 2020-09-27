package filter

// Prime filters prime numbers
type Prime struct{}

// NewPrime creates a prime-finding Filter.
func NewPrime() Filter {
	return &Prime{}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (p *Prime) Run(in <-chan int, out chan<- int) {
	defer close(out)

	for v := range in {
		if IsPrime(v) {
			out <- v
		}
	}
}

func IsPrime(n int) bool {
	if n <= 3 {
		return n > 1
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}
