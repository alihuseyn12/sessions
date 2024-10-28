package chanels_basic_operations

import "time"

func UnbufferdChanalsCheck() chan int {

	ch := make(chan int)

	go func() {
		time.Sleep(time.Second * 5)
		ch <- 42
	}()
	return ch
}
