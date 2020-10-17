package filter

// cumSummer sums all received values cumulative.
type cumSummer struct{}

// NewCumSummer creates a new cumulative sum Filter
func NewCumSummer() Filter {
	return &cumSummer{}
}

// Run against an input and output channel until the input channel is closed.
// The output channel is closed on return.
func (cs *cumSummer) Run(in <-chan int, out chan<- int) {
	defer close(out)

	cumSum := 0
	for v := range in {
		cumSum += v
		out <- cumSum
	}
}
