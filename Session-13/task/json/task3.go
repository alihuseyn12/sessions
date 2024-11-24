package json

import (
	"encoding/json"
	"fmt"
)

type Student []struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

func PassStudents(f []byte) {

	var data Student

	err1 := json.Unmarshal(f, &data)
	if err1 != nil {
		fmt.Println(err1)

	}
	fmt.Println("Students with passing grades:")
	for i := 0; i < len(data); i++ {
		if data[i].Grade < 60 {
			fmt.Println(data[i].Name)
		}
	}
}
