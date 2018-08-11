package shapes

import (
	"math"
)

//Rectangle test
type Rectangle struct {
	Width  float64
	Height float64
}

//GetAera return aera
func (r Rectangle) GetAera() float64 {
	return r.Height * r.Width
}

//Circle struct
type Circle struct {
	Radius float64
}

//GetAera return aera
func (r Circle) GetAera() float64 {
	return math.Pi * r.Radius * r.Radius
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
