package goroutines

import (
	"fmt"
	"time"
)

func CheckProcess() {
	alfabit := [5]string{"A", "B", "C", "D", "E"}
	go func() {
		for i := 0; i < len(alfabit); i++ {
			fmt.Println(alfabit[i])
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println(i)
			time.Sleep(time.Second * 3)
		}
	}()
	time.Sleep(time.Second * 11)
}
