//Explanation:
//Shape Interface: The base for all shapes, providing the Accept method to allow visitors to operate on them.
//Concrete Shapes (Circle, Rectangle): Implement the Shape interface and include the logic for accepting a visitor.
//Visitor Interface (ShapeVisitor): Declares methods for visiting different shape types.
//Concrete Visitor (AreaCalculator): Implements the visitor interface and performs specific actions (like area calculation) on different shapes.
//ShapeCollection: Uses the visitor by calling the Accept method for all shapes in the collection.

package main

import "fmt"

// Shape is the interface for all shapes
type Shape interface {
	Accept(visitor ShapeVisitor)
}

// Circle represents a circle shape
type Circle struct {
	radius float64
}

// Rectangle represents a rectangle shape
type Rectangle struct {
	width, height float64
}

// Visitor interface defines a method for each concrete shape
type ShapeVisitor interface {
	VisitCircle(circle *Circle)
	VisitRectangle(rectangle *Rectangle)
}

// AreaCalculator calculates areas of shapes
type AreaCalculator struct {
	totalArea float64
}

func (a *AreaCalculator) VisitCircle(circle *Circle) {
	area := 3.14 * circle.radius * circle.radius
	fmt.Printf("Area of Circle: %.2f\n", area)
	a.totalArea += area
}

func (a *AreaCalculator) VisitRectangle(rectangle *Rectangle) {
	area := rectangle.width * rectangle.height
	fmt.Printf("Area of Rectangle: %.2f\n", area)
	a.totalArea += area
}

// ShapeCollection is a collection of shapes
type ShapeCollection struct {
	shapes []Shape
}

// Accept allows the visitor to access the shapes in the collection
func (s *ShapeCollection) Accept(visitor ShapeVisitor) {
	for _, shape := range s.shapes {
		shape.Accept(visitor)
	}
}

// Accept method for Circle
func (c *Circle) Accept(visitor ShapeVisitor) {
	visitor.VisitCircle(c)
}

// Accept method for Rectangle
func (r *Rectangle) Accept(visitor ShapeVisitor) {
	visitor.VisitRectangle(r)
}

func main() {
	circle := &Circle{radius: 5}
	rectangle := &Rectangle{width: 4, height: 6}

	shapes := &ShapeCollection{
		shapes: []Shape{circle, rectangle},
	}

	areaCalculator := &AreaCalculator{}
	shapes.Accept(areaCalculator)

	fmt.Printf("Total Area: %.2f\n", areaCalculator.totalArea)
}

//OUTPUT:
//
//Area of Circle: 78.50
//Area of Rectangle: 24.00
//Total Area: 102.50
