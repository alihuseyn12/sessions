package defining

type Speaker interface {
	Speak() string
}
type Dog struct{}
type Person struct{}

func (d Dog) Speak() string {

	return "Dog says: woof"
}
func (p Person) Speak() string {
	return "Person says: Hello!"
}
