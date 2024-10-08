package main

import "fmt"

func incrementByValue(x int) int {
	return x + 1
}

func incrementByPointer(x int) *int {

	ptr := &x
	*ptr = x + 1
	return ptr

}

func main() {
	var x int = 5
	fmt.Printf("Value of x before incrementing by value:= %d \n", x)
	val := incrementByValue(x)
	fmt.Printf("Value of x after incrementing by value:= %d \n", val)
	fmt.Printf("Value of x before incrementing by pointer:= %d \n", x)
	ptr := incrementByPointer(x)
	fmt.Printf("Value of x before incrementing by pointer:= %d \n", *ptr)

}
