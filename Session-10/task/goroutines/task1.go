package goroutines

import (
	"fmt"
	"time"
)

func CountNumber() {

	fmt.Println("Main function started")
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	time.Sleep(time.Second * 1)
	fmt.Println("Main function ended")

}
