package xml

import (
	"encoding/xml"
	"fmt"
)

type Employee struct {
	Name     string `xml:"name"`
	Position string `xml:"position"`
	Salary   int    `xml:"salary"`
}
type Employees struct {
	Employee []Employee `xml:"employee"`
}

func AboveSalary(f []byte) {
	var employees Employees
	err1 := xml.Unmarshal(f, &employees)
	if err1 != nil {
		fmt.Println("Unmarshal err:", err1)
	}
	for _, emp := range employees.Employee {
		if emp.Salary > 50000 {
			fmt.Println("Employees with Salary above 50000:")
			fmt.Printf("-%s,%s", emp.Name, emp.Position)
		}

	}
}
