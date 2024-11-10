package buffered_unbufered_channels

import (
	"fmt"
	"time"
)

func CheckBufferedChannel(ch chan int) {

	go func() {

		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}()
	close(ch)
	time.Sleep(time.Second * 1)
	fmt.Println("All values received")
}
