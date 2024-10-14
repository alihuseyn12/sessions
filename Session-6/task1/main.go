package main

import (
	"fmt"
)

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	var sum int
	fmt.Printf("First element: %d\n", arr[0])
	fmt.Printf("Last element: %d\n", arr[len(arr)-1])
	for _, v := range arr {
		sum += v
	}
	fmt.Printf("Sum: %d\n", sum)
	var j int
	for i := len(arr) - 1; i > j; i-- {
		arr[j], arr[i] = arr[i], arr[j]
		j++

	}
	fmt.Printf("Reversed array: %d", arr)
	//task 1 is done

}
