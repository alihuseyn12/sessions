package sync_waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func Task3(a int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Goroutine %d starting\n", a)
	time.Sleep(time.Second * 1)
	fmt.Printf("Goroutine %d finished\n", a)

}
