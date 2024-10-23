package implementing

import "fmt"

//Task 4: Using Interfaces with Functions

type Notifier interface {
	Notify(message string)
}

// and SMSNotifier
type EmailNotifier struct{}
type SmsNotifier struct{}

func (e EmailNotifier) Notify(message string) {
	fmt.Println("Sending email notification:", message)
}
func (s SmsNotifier) Notify(message string) {
	fmt.Println("Sending SMS notification:", message)
}
