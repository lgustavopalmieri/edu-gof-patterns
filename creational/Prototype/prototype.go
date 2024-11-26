//Explanation:
//Prototype Interface (Shape): Declares the Clone method for cloning prototype objects.
//Concrete Prototypes (Rectangle, Circle): Implement the Shape interface, providing their own Clone methods that return a copy of themselves.
//Cloning: Each clone method in concrete prototypes constructs a new instance with copied field values, simulating a shallow copy.
//Client Code: Creates instances of Rectangle and Circle, then clones these objects using the Clone method, working with prototypes rather than directly creating new instances.

package main

import (
	"fmt"
)

// Shape is the prototype interface with a Clone method
type Shape interface {
	Clone() Shape
	GetType() string
}

// Rectangle is a concrete prototype
type Rectangle struct {
	Width  int
	Height int
}

// Clone creates a copy of this Rectangle
func (r *Rectangle) Clone() Shape {
	return &Rectangle{
		Width:  r.Width,
		Height: r.Height,
	}
}

// GetType returns the type of shape
func (r *Rectangle) GetType() string {
	return "Rectangle"
}

// Circle is another concrete prototype
type Circle struct {
	Radius int
}

// Clone creates a copy of this Circle
func (c *Circle) Clone() Shape {
	return &Circle{
		Radius: c.Radius,
	}
}

// GetType returns the type of shape
func (c *Circle) GetType() string {
	return "Circle"
}

func main() {
	// Create initial set of objects
	rectangle := &Rectangle{Width: 10, Height: 20}
	circle := &Circle{Radius: 15}

	// Clone the objects
	rectangleClone := rectangle.Clone()
	circleClone := circle.Clone()

	fmt.Printf("Original: %s with dimensions %d x %d\n", rectangle.GetType(), rectangle.Width, rectangle.Height)
	fmt.Printf("Clone: %s with dimensions %d x %d\n", rectangleClone.GetType(), rectangleClone.(*Rectangle).Width, rectangleClone.(*Rectangle).Height)

	fmt.Printf("Original: %s with radius %d\n", circle.GetType(), circle.Radius)
	fmt.Printf("Clone: %s with radius %d\n", circleClone.GetType(), circleClone.(*Circle).Radius)
}
