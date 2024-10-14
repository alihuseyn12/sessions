package main

import "fmt"

func main() {

	//res := reflect.TypeOf(cauntry["Azerbaijan"])
	//fmt.Println(res)
	a := countryCek("France ")
	println(a)

}
func countryCek(name string) string {

	country := map[string]string{
		"Azerbaijan": "Baku",
		"USA":        "Washington",
		"Germany":    "Berlin",
		"Japan":      "Tokyo",
		"Turkey ":    "Ankara",
		"Italy ":     "Rome",
	}
	for k, v := range country {
		fmt.Printf("Capital of %s is %s \n", k, v)
	}
	if value, ok := country[name]; ok {
		return value
	} else {
		return "Capital of " + name + "is not found"
	}

}
