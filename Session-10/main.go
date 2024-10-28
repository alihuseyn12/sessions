package main

import "Session-10/task/buffered_unbufered_channels"

func main() {
	//task1
	//goroutines.CountNumber()
	//time.Sleep(1 * time.Second)

	//task2

	//goroutines.CheckProcess()
	//time.Sleep(time.Second * 2)
	//fmt.Println("Main finished")

	//task3 Channels - Basic Operations

	//ch := chanels_basic_operations.UnbufferdChanalsCheck()
	//data := <-ch
	//fmt.Println("Received value:", data)

	//task4

	//ch1 := chanels_basic_operations.ChanlTask4()
	//for ch2 := range ch1 {
	//	fmt.Println(ch2)
	//}
	//fmt.Println("Channel closed")

	//task5

	//buffered_unbufered_channels.CheckBufferedChannel(10, 20, 30)

	//task6

	buffered_unbufered_channels.UnbufferedchannelChek6("hello")

}
