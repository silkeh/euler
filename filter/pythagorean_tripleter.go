package filter

// pythagoreanTripleter filters palindrome numbers.
type pythagoreanTripleter struct {
	n int
}

// NewPythagoreanTripleter creates a palindrome-finding Filter.
func NewPythagoreanTripleter(n int) Filter {
	return &pythagoreanTripleter{n: n}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (pt *pythagoreanTripleter) Run(in <-chan int, out chan<- int) {
	defer close(out)
	for {
		m, ok1 := <-in
		n, ok2 := <-in
		if !(ok1 && ok2) {
			return
		}
		if m < n {
			n, m = m, n
		}
		a, b, c := pt.generateTriplet(m, n)
		if pt.isTriplet(a, b, c) {
			for k := 1; k <= pt.n; k++ {
				out <- a * k
				out <- b * k
				out <- c * k
			}
		}
	}
}

// isTriplet tests if a, b, c, is a Pythagorean Triplet
func (pt *pythagoreanTripleter) isTriplet(a, b, c int) bool {
	return a*a+b*b == c*c
}

// generateTriplet creates a Pythagorean Triplet from m and n
// m and n must be coprime.
func (pt *pythagoreanTripleter) generateTriplet(m, n int) (a, b, c int) {
	m2 := m * m
	n2 := n * n
	a, b, c = m2-n2, 2*m*n, m2+n2
	if m%2 == 1 && n%2 == 1 {
		a, b, c = a/2, b/2, c/2
	}
	return a, b, c
}
