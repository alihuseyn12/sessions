package buffered_unbufered_channels

import (
	"fmt"
	"time"
)

func CheckBufferedChannel(a, b, c int) {
	ch := make(chan int, 3)
	ch <- a
	ch <- b
	ch <- b

	close(ch)
	fmt.Println("Sent values into buffered channel")
	go func() {

		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}()

	time.Sleep(time.Second * 1)
	fmt.Println("All values received")
}
