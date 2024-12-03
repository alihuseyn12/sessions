package using_reflection

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {

	return "Hello, I am " + p.Name + "."
}

func InvokeMethod(value interface{}, methodName string) {

	//typ := reflect.TypeOf(&value)
	val := reflect.ValueOf(value).Elem()
	val.FieldByName("Name").SetString("Alice")
	res := val.MethodByName(methodName)

	//Invoked method: Greet, Result: Hello, I am Alice.
	fmt.Println("Invoked method:", methodName, ",", "Result:", res.Call(nil))
}
