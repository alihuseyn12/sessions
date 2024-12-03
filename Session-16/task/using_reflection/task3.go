package using_reflection

import (
	"reflect"
)

func SetFieldValue(value interface{}, fieldName string, newValue string) {
	val := reflect.ValueOf(value).Elem()
	val.FieldByName(fieldName).SetString(newValue)
}
