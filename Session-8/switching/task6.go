package switching

import "fmt"

//Task 6: Type Switch
//
//Write a program that takes an interface{} and uses a type switch to determine whether the underlying value is an int, string, or bool.
//Print a specific message for each type.
//If the type is not recognized, print "Unknown type".
//Expected Output:
//
//Type is int: 100
//Type is string: Hi, Gophers
//Type is bool: true
//Unknown type

func CheckContentType(content any) {
	switch v := content.(type) {
	case string:

		fmt.Printf("Type is %T:  %s\n", v, v)
	case int:

		fmt.Printf("Type is %T, %d\n", v, v)
	case bool:
		fmt.Printf("Type is %T, %t\n", v, v)
	default:
		fmt.Println("Type is Unknown")

	}
}
