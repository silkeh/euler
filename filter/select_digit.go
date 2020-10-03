package filter

import "strconv"

// SelectDigit selects a digit from a given grid of values
type SelectDigit struct {
	givenDigits []string
	n           int
}

// NewSelectDigit creates a SelectDigit filter
func NewSelectDigit(digit []string, n int) Filter {
	return &SelectDigit{givenDigits: digit, n: n}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (sd *SelectDigit) Run(in <-chan int, out chan<- int) {
	defer close(out)

	for {
		x, ok1 := <-in
		y, ok2 := <-in
		if !(ok1 && ok2) {
			return
		}
		a, err := strconv.Atoi(sd.givenDigits[y][x*sd.n : (x+1)*sd.n])
		if err != nil {
			return
		}
		out <- a
	}

}
