package main

import "fmt"

func main() {

	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice := a[2:6]
	fmt.Printf("Sub-slice: %d\n", slice)
	slice = append(slice, 10, 11, 12)
	fmt.Printf("Append-slice: %d\n", slice)
	fmt.Printf("Print original array: %d\n", a)

	// Show the effect of append() on the underlying array by printing both the slice and the original array after appending.
	//yuxaridaki tapsiriqi tam olaraq basa dusmedim.

}
