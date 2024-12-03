package reflection

import (
	"fmt"
	"reflect"
)

func IdentifyTypeAndKind(value interface{}) {
	typ := reflect.TypeOf(value)
	val := reflect.ValueOf(value)
	fmt.Printf(" Val: %v,Type: %v,Kind: %v \n", val, typ, typ.Kind())

}
