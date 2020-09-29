package generator

type DigitReader struct {
	givenDigits []byte
}

// NewDigitReader creates a DigitReader Generator.
func NewDigitReader(digit []byte) Generator {
	return &DigitReader{givenDigits: digit}
}

// Run against a channel.
// The channel is closed when the DigitReader is done.
func (dr *DigitReader) Run(out chan<- int, done <-chan bool) {
	defer close(out)
	for i := 0; i < len(dr.givenDigits); i++ {
		out <- int(dr.givenDigits[i]) - 0x30
	}
}
