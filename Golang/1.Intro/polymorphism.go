/*
Polymorphism is a fundamental concept in object-oriented programming (OOP) that
allows objects of different classes to be treated as if they were objects of
the same class.
It is the ability of an object to take on many forms or have multiple behaviors,
depending on the context in which it is used.

Polymorphism is achieved through the use of inheritance and interfaces.
Inheritance allows a subclass to inherit the methods and properties of its
superclass, while interfaces define a set of methods that a class must implement.
By using polymorphism, a program can be written to handle objects of different
classes in a uniform way, without having to know the specific class of each
object at runtime.
*/

package main

import (
	"fmt"
	"math"
)

// Define a Shape interface with methods for calculating the area and perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a Circle struct that implements the Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Round((math.Pi*c.Radius*c.Radius)*100) / 100
}

func (c Circle) Perimeter() float64 {
	return math.Round((2*math.Pi*c.Radius)*100) / 100
}

// Define a Rectangle struct that implements the Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return math.Round((r.Width*r.Height)*100) / 100
}

func (r Rectangle) Perimeter() float64 {
	return math.Round((2*r.Width+2*r.Height)*100) / 100
}

// Define a function that takes a Shape as input and prints its area and perimeter
func PrintShapeInfo(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
	// Create a Circle and a Rectangle
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 5}

	// Call the PrintShapeInfo function with each shape
	fmt.Println("Circle:")
	PrintShapeInfo(c)

	fmt.Println("Rectangle:")
	PrintShapeInfo(r)
}

// Output:
/*
Circle:
Area: 78.54
Perimeter: 31.42
Rectangle:
Area: 50
Perimeter: 30
*/
