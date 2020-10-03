package generator

// DirectionGenerator generates a sequence of numbers.
type DirectionGenerator struct {
	x, y, n int
}

// NewDirectionGenerator creates a Sequence Generator.
func NewDirectionGenerator(x, y, n int) Generator {
	return &DirectionGenerator{x: x, y: y, n: n}
}

// Run against a channel.
// The channel is closed when the DirectionGenerator is done.
func (dg *DirectionGenerator) Run(out chan<- int, done <-chan bool) {
	defer close(out)
	diagx, diagy := dg.x-dg.n, dg.y-dg.n
	horx, hory := dg.x-dg.n, dg.y
	verx, very := dg.x, dg.y-dg.n
	for x := 0; x < dg.x; x++ {
		for y := 0; y < dg.y; y++ {

			//diagonal
			if x <= diagx && y <= diagy {
				dg.diagonalDownGenerator(out, x, y)
				dg.diagonalUpGenerator(out, x, y)
			}

			if x <= diagx && dg.n < y && y <= diagy {
			}

			//Horizontal
			if x <= horx && y <= hory {
				dg.horizontalGenerator(out, x, y)
			}

			//Vertical
			if x <= verx && y <= very {
				dg.verticalGenerator(out, x, y)
			}
		}
	}
}

// diagonalUpGenerator outputs a series of x, y positions
func (dg *DirectionGenerator) diagonalUpGenerator(out chan<- int, i, j int) {
	for k := 0; k < dg.n; k++ {
		out <- i + k
		out <- j + dg.n - 1 - k
	}
}

// diagonalDownGenerator outputs a series of x, y positions
func (dg *DirectionGenerator) diagonalDownGenerator(out chan<- int, i, j int) {
	for k := 0; k < dg.n; k++ {
		out <- i + k
		out <- j + k
	}
}

// horizontalGenerator outputs a series of x, y positions
func (dg *DirectionGenerator) horizontalGenerator(out chan<- int, i, j int) {
	for k := 0; k < dg.n; k++ {
		out <- i + k
		out <- j
	}
}

// verticalGenerator outputs a series of x, y positions
func (dg *DirectionGenerator) verticalGenerator(out chan<- int, i, j int) {
	for k := 0; k < dg.n; k++ {
		out <- i
		out <- j + k
	}
}
