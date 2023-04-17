package geometry

import "math"

type Rectangle struct {
	Height float64
	Width float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height float64
	Base float64
}

type Shape interface {
	CalcArea() float64
}

// methods have a receiver, and that gives us access to a reference of the data on the receiver variable
// any type that satisfies the interface will compile, in this case Rectangle and Circle both have CalcArea methods so they satisfy the interface
func (r Rectangle) CalcArea() float64 {
	return r.Height * r.Width
}

func (c Circle) CalcArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) CalcArea() float64 {
	return t.Height * t.Base / 2
}

func calcPerimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

// func calcArea(rectangle Rectangle) float64 {
// 	return rectangle.Height * rectangle.Width
// }