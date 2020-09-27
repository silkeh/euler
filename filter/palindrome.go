package filter

import "strconv"

// Palindrome filters palindrome numbers.
type Palindrome struct{}

// NewPalindrome creates a palindrome-finding Filter.
func NewPalindrome() Filter {
	return &Palindrome{}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (f *Palindrome) Run(in <-chan int, out chan<- int) {
	defer close(out)

	for v := range in {
		if isPalindrome(v) {
			out <- v
		}
	}
}

func isPalindrome(n int) bool {
	s := []rune(strconv.Itoa(n))
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
