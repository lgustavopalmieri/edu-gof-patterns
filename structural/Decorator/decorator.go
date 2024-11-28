//Explanation:
//Component Interface: Component defines the Execute method that both concrete components and decorators implement.
//ConcreteComponent: Implements the base behavior of the Component.
//Decorator: Holds a reference to a Component object and implements the same interface. It intercepts Execute calls to add additional behavior.
//FancyDecorator: Further decorates the Decorator, showing how multiple decorators can be layered.
//Client Code: Composes decorators with components, allowing flexible addition of responsibilities at runtime.

package main

import (
	"fmt"
)

// Component is the interface for objects that can have responsibilities added dynamically
type Component interface {
	Execute() string
}

// ConcreteComponent is a concrete implementation of Component
type ConcreteComponent struct{}

// Execute returns the base functionality
func (c *ConcreteComponent) Execute() string {
	return "Hello"
}

// Decorator provides additional functionality while conforming to Component interface
type Decorator struct {
	component Component
}

// Execute calls the underlying component and adds extra behavior
func (d *Decorator) Execute() string {
	return fmt.Sprintf("%s, World!", d.component.Execute())
}

// FancyDecorator adds even more behavior on top
type FancyDecorator struct {
	component Component
}

// Execute calls the underlying component and adds its own behavior
func (f *FancyDecorator) Execute() string {
	return fmt.Sprintf("*** %s ***", f.component.Execute())
}

func main() {
	// Create the base component
	component := &ConcreteComponent{}

	// Decorate the component with Decorator
	decorator := &Decorator{component: component}

	// Further decorate with FancyDecorator
	fancyDecorator := &FancyDecorator{component: decorator}

	// Execute the final decorated component
	fmt.Println(fancyDecorator.Execute())
}
