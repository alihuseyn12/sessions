package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a

}
func main() {
	x := 10
	y := 20
	//Before swapping: x = 10, y = 20
	fmt.Printf("Before swapping x=%d: y=%d \n", x, y)
	swap(&x, &y)
	fmt.Printf("After swapping x=%d: y=%d \n", x, y)

}
