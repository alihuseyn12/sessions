package main

import "Session-16/task/using_reflection"

func main() {
	////Task1
	//reflection.IdentifyTypeAndKind(42)
	//reflection.IdentifyTypeAndKind("Hello")
	//reflection.IdentifyTypeAndKind([]int{1, 2, 3})

	//Task2
	//type Person struct {
	//	Name string
	//	Age  int
	//}
	//p := &Person{Name: "Ali", Age: 30}
	//reflection.InspectStruct(p)

	////Task3
	//type Person struct {
	//	Name string
	//	Age  int
	//	City string
	//}
	//p := &Person{Name: "Ali", Age: 30, City: "Baku"}
	//using_reflection.SetFieldValue(p, "City", "Ganja")
	//fmt.Printf("Name:%s,Age:%d,City:%s\n", p.Name, p.Age, p.City)

	////task4

	p := &using_reflection.Person{Name: "Ali", Age: 30}
	using_reflection.InvokeMethod(p, "Greet")
	////Task5

	//fmt.Println(go_generics.Sum([]int{1, 2, 3, 4}))
	//fmt.Println(go_generics.Sum([]float64{1.5, 2.5, 3.5}))

	//// task6
	//min, max := go_generics.MinMax([]int{10, 20, 5, 8, 15})
	//fmt.Println(min, max)

}
