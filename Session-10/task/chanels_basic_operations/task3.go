package chanels_basic_operations

import "time"

func UnbufferdChanalsCheck() chan int {

	ch := make(chan int)

	go func() {
		time.Sleep(time.Millisecond * 500)
		ch <- 42
	}()
	return ch
}
