//Explanation:
//Component Interface: Component defines the common interface with methods like Draw that both individual and composite objects implement.
//Leaf Nodes (Circle, Rectangle): These are individual objects that directly implement the Component interface.
//Composite Node (Composite): Contains child components, which could be other composite nodes or leaf nodes, and implements the Component interface.
//Client Code: Uses the Composite and Leaf classes interchangeably. In main, it configures a composite structure (group) and invokes operations on it, leading to uniform handling of complex compositions and simple elements.

package main

import "fmt"

// Component is the interface for all objects in the composition
type Component interface {
	Draw()
}

// Circle is a leaf component
type Circle struct {
	Name string
}

// Draw method for Circle
func (c *Circle) Draw() {
	fmt.Printf("Drawing Circle: %s\n", c.Name)
}

// Rectangle is another leaf component
type Rectangle struct {
	Name string
}

// Draw method for Rectangle
func (r *Rectangle) Draw() {
	fmt.Printf("Drawing Rectangle: %s\n", r.Name)
}

// Composite is a composite component that can hold children
type Composite struct {
	Children []Component
	Name     string
}

// Draw method for Composite
func (composite *Composite) Draw() {
	fmt.Printf("Drawing Composite: %s\n", composite.Name)
	for _, child := range composite.Children {
		child.Draw()
	}
}

func (composite *Composite) Add(component Component) {
	composite.Children = append(composite.Children, component)
}

func main() {
	circle1 := &Circle{Name: "Circle-1"}
	circle2 := &Circle{Name: "Circle-2"}
	rectangle := &Rectangle{Name: "Rectangle-1"}

	// Create a composite and add the leaf components
	group := &Composite{Name: "Group-1"}
	group.Add(circle1)
	group.Add(circle2)
	group.Add(rectangle)

	// Draw the composite, which in turn draws its children
	group.Draw()
}
