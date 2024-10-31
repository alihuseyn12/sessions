package main

import (
	"Session-9/task/wraperr"
	"errors"
	"fmt"
)

func main() {
	//task1
	//res, err := custerr.Divide(25, 0)
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("Result: %f", res)
	//}

	//task3
	openFile := wraperr.OpenFile("nonexistent.txt")
	fmt.Printf("Wrapped Error:%v\n ", openFile)

	res := errors.Unwrap(openFile)
	fmt.Printf("Orginal Error", res)

}
