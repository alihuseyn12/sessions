package reflection

import (
	"fmt"
	"reflect"
)

func InspectStruct(value interface{}) {
	typ := reflect.TypeOf(value).Elem()
	val := reflect.ValueOf(value).Elem()

	val.FieldByName("Name").SetString("Alice")
	fmt.Println()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Name == "Name" {
			fmt.Printf("Field Name:%s,Type:%s,Value:%s\n", field.Name, field.Type, val.Field(i).String())
		} else {
			fmt.Printf("Field Name:%s,Type:%s,Value:%d\n", field.Name, field.Type, val.Field(i).Int())
		}

		// Field Name: Name, Type: string, Value: Alice
	}
}
