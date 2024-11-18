package main

import (
	"Session-13/task/xml"
	"fmt"
	"os"
)

func main() {
	// Task1
	//file, err := os.OpenFile("task/data.csv", os.O_RDONLY, 0)
	//defer file.Close()
	//if err != nil {
	//	fmt.Println("error is :", err)
	//
	//}
	//reading_writing.PassOfStudent(file)

	//Task2
	//
	//f, err := os.OpenFile("task/story.txt", os.O_RDONLY, 0)
	//if err != nil {
	//	fmt.Println("err is:", err)
	//	return
	//}
	//defer f.Close()
	//reading_writing.LineCount(f)

	//Task3
	//
	//f, err := os.ReadFile("task/students.json")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//json2.PassStudents(f)

	//Task4

	//f, err := os.ReadFile("task/config.json")
	//if err != nil {
	//	fmt.Println("ReadFile err:", err)
	//}
	//json2.CheckConfigFile(f)

	//Task5

	//f, err := os.ReadFile("task/employees.xml")
	//if err != nil {
	//	fmt.Println("ReadFile err:", err)
	//}
	//xml.AboveSalary(f)

	//Task6

	f, err := os.OpenFile("task/config.xml", os.O_CREATE|os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	xml.EncodeFile(f)

}
