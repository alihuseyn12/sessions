package main

import "fmt"

func main() {
	var x int = 10
	ptr := &x
	fmt.Println("Value of x=", x)
	fmt.Println("Address of x=", &ptr)
	fmt.Println("Address of ptr=", *ptr)

	//fmt.Println("Value of x=", x)
	//fmt.Println("Address of x=", &ptr)
	//fmt.Println("Address of ptr=", *ptr)
	//fmt.Println("Value of x=", x)
	//fmt.Println("Address of x=", &ptr)
	//fmt.Println("Address of ptr=", *ptr)
}
