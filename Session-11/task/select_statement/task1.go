package select_statement

import "fmt"

func ExampleSelectStatment1(ch1, ch2 chan string) {

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)

	}
}
