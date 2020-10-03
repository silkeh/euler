package generator

import "strconv"

// DigitReader is a Generator that sends a byte slice
// containing ASCII numbers in sequence.
type DigitReader struct {
	givenDigits string
	n           int
}

// NewDigitReader creates a DigitReader Generator.
func NewDigitReader(digit string, n int) Generator {
	return &DigitReader{givenDigits: digit, n: n}
}

// Run against a channel.
// The channel is closed when the DigitReader is done.
func (dr *DigitReader) Run(out chan<- int, done <-chan bool) {
	defer close(out)
	for i := 0; i < len(dr.givenDigits); i += dr.n {
		a, err := strconv.Atoi(dr.givenDigits[i : i+dr.n])
		if err != nil {
			return
		}
		out <- a
	}
}
