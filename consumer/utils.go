package consumer

// channelCleanup will send the signal
func channelCleanup(in <-chan int, done chan<- bool) {
	close(done)
	for range in {
		// do nothing
	}
}
