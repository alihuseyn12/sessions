package main

import (
	"fmt"
	"session-8/assertion"
	"session-8/defining"
	"session-8/embedding"
	"session-8/implementing"
	"session-8/switching"
)

func main() {
	fmt.Println("----//task1-----")
	var circle defining.Shape
	var rectangle defining.Shape
	circle = defining.Cricle{
		Radius: 5,
	}
	rectangle = defining.Rectangle{
		Height: 10,
		Width:  5,
	}

	fmt.Println("Circle area:", circle.Area())
	fmt.Println("Rectangle area:", rectangle.Area())

	fmt.Println("----//task2-----")
	var person defining.Speaker
	var dog defining.Speaker
	person = defining.Person{}
	fmt.Println(person.Speak())
	dog = defining.Dog{}
	fmt.Println(dog.Speak())

	fmt.Println("----//task3-----")

	var ciditCart implementing.PaymentProcessor
	var payPal implementing.PaymentProcessor
	payPal = implementing.PayPal{}
	ciditCart = implementing.CreditCard{}
	ciditCart.ProcessPayment(100)
	payPal.ProcessPayment(75.5)

	fmt.Println("----//task4-----")

	var emailNotifier implementing.Notifier
	emailNotifier = implementing.EmailNotifier{}
	emailNotifier.Notify("Your item has shipped")
	var smsNotifier implementing.Notifier
	smsNotifier = implementing.SmsNotifier{}
	smsNotifier.Notify("Your item has been shipped")

	fmt.Println("----//task5-----")

	var number assertion.EmptyInterface
	number = 42
	value, ok := number.(int)
	if ok {
		fmt.Printf("Value is of type %T :%d \n", value, value)
	}
	number = "golang"
	fmt.Printf("%T", number)
	var stringT assertion.EmptyInterface
	stringT = "Golang"
	value1, ok1 := stringT.(string)
	if ok1 {
		fmt.Printf("Value is of type %T :%s \n", value1, value1)
	}
	var floatT assertion.EmptyInterface
	floatT = 3.14
	value2, ok2 := floatT.(float64)
	if ok2 {
		fmt.Printf("Value is of type %T :%f \n", value2, value2)
	}

	fmt.Println("----//task6-----")

	switching.CheckContentType(100)
	switching.CheckContentType("Hi,Gophers")
	switching.CheckContentType(true)

	fmt.Println("----//task7-----")
	var someThink embedding.Robot
	someThink = embedding.AutoBot{}
	someThink.Say()
	someThink.Move()

}
