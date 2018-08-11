package shapes

import (
	"math"
)

//Shape interface Area
type Shape interface {
	Area() float64
}

//Rectangle struct Width Height
type Rectangle struct {
	Width  float64
	Height float64
}

//Area return area
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

//Circle struct Radius
type Circle struct {
	Radius float64
}

//Area return area
func (r Circle) Area() float64 {
	return math.Pi * r.Radius * r.Radius
}

//Triangle struct Base Height
type Triangle struct {
	Base   float64
	Height float64
}

//Area return area
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

// Perimeter return
func Perimeter(rectangle Rectangle) float64 {
	pm := (rectangle.Height + rectangle.Width) * 2
	return pm
}

// Area return square
func Area(rectangle Rectangle) float64 {
	sq := rectangle.Height * rectangle.Width
	return sq
}
