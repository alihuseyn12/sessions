package defining

import "math"

//Define an interface Shape with the method Area() float64. Then, implement this interface for two structs: Circle and Rectangle.
//The Circle struct should have a field for radius.
//The Rectangle struct should have fields for width and height.
//Write a function that takes a Shape and prints the area of that shape.

//Expected Output:
//
//Circle area: 78.53981633974483
//Rectangle area: 50

type Shape interface {
	Area() float64
}
type Cricle struct {
	Radius float64
}
type Rectangle struct {
	Height float64
	Width  float64
}

func (s Cricle) Area() float64 {
	return math.Pi * s.Radius * s.Radius
}
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
