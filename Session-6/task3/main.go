package main

import "fmt"

func main() {
	an := make([]int, 3, 5)
	fmt.Printf("len is = %d cap is = %d\n", len(an), cap(an))
	an = []int{1, 2, 3, 4, 5}
	fmt.Printf("len is = %d cap is = %d\n", len(an), cap(an))
	//task 3 is done
}
