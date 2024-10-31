package main

import (
	"fmt"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		//	time.Sleep(time.Second * 2)
		ch1 <- "frist canal1"
	}()

	go func() {
		//time.Sleep(time.Second * 1)
		ch2 <- "secnd canal2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)

		}
	}
}
