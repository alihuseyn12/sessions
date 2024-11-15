package rate_limit_and_throttling

import (
	"fmt"
	"time"
)

func CreateTaskExecuted(ticker *time.Ticker) {

	for i := 0; i < 5; i++ {
		select {
		case <-ticker.C:
			fmt.Println("Task executed")

		}

	}
	fmt.Println("(ticker stopped)")

}
