package main

import "fmt"

func doubleVariadicElements(numbers ...*int) {
	for _, num := range numbers {
		*num = *num * 2

	}
	for _, num := range numbers {
		fmt.Printf("%d ", *num)
	}
}

func main() {
	a := 1
	b := 2
	c := 3

	doubleVariadicElements(&a, &b, &c)

}
