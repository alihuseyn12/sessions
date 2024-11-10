package chanels_basic_operations

func ChanlTask4(ch chan int) {

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		close(ch)

	}()

}
