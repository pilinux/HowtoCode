// A `struct` is a named collection of fields to store data
// Method - func (receiverName recieverType) methodName(args if any) returnType {}
package shapes

import "math" // To use `Pi`

type rectangle struct {
	height float64
	width  float64
}

type circle struct {
	radius float64
}

type triangle struct {
	base   float64
	height float64
}

// Methods
// perimeter of a rectangle
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// area of a rectangle
func (r rectangle) area() float64 {
	return r.width * r.height
}

// perimeter or circumference of a circle
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// area of a circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// area of a triangle
func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}
