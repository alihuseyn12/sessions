package buffered_unbufered_channels

func UnbufferedchannelChek6(s string) {
	ch := make(chan string)
	ch <- s
	go func() {

	}()
}
