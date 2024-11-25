//Explanation:
//Product Interfaces (Button, Checkbox): Define the operations that all concrete product variants must implement.
//Concrete Products (MacOSButton, WindowsButton, MacOSCheckbox, WindowsCheckbox): Different implementations of the product interfaces for each operating system.
//Factory Interface (WidgetFactory): Declares methods for creating each of the abstract products.
//Concrete Factories (MacOSFactory, WindowsFactory): Implement the abstract factory interface and produce a family of products that work together.
//Client Code: Uses GetFactory to obtain a WidgetFactory, then uses it to create and manipulate product objects without being aware of their concrete classes.
//This pattern is particularly useful when a system needs to be independent of how its objects were created, composed, and represented.

package main

import "fmt"

// Button represents the Button product interface
type Button interface {
	Paint() string
}

// Checkbox represents the Checkbox product interface
type Checkbox interface {
	Paint() string
}

// MacOSButton concrete product
type MacOSButton struct{}

// Paint implementation for MacOS button
func (m *MacOSButton) Paint() string {
	return "Painting MacOS Button"
}

// WindowsButton concrete product
type WindowsButton struct{}

// Paint implementation for Windows button
func (w *WindowsButton) Paint() string {
	return "Painting Windows Button"
}

// MacOSCheckbox concrete product
type MacOSCheckbox struct{}

// Paint implementation for MacOS checkbox
func (m *MacOSCheckbox) Paint() string {
	return "Painting MacOS Checkbox"
}

// WindowsCheckbox concrete product
type WindowsCheckbox struct{}

// Paint implementation for Windows checkbox
func (w *WindowsCheckbox) Paint() string {
	return "Painting Windows Checkbox"
}

// WidgetFactory interface for creating families of related objects
type WidgetFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// MacOSFactory concrete factory for MacOS products
type MacOSFactory struct{}

// CreateButton creates a MacOS button
func (m *MacOSFactory) CreateButton() Button {
	return &MacOSButton{}
}

// CreateCheckbox creates a MacOS checkbox
func (m *MacOSFactory) CreateCheckbox() Checkbox {
	return &MacOSCheckbox{}
}

// WindowsFactory concrete factory for Windows products
type WindowsFactory struct{}

// CreateButton creates a Windows button
func (w *WindowsFactory) CreateButton() Button {
	return &WindowsButton{}
}

// CreateCheckbox creates a Windows checkbox
func (w *WindowsFactory) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

// GetFactory returns a WidgetFactory based on input
func GetFactory(osType string) (WidgetFactory, error) {
	switch osType {
	case "mac":
		return &MacOSFactory{}, nil
	case "windows":
		return &WindowsFactory{}, nil
	default:
		return nil, fmt.Errorf("unknown OS type")
	}
}

func main() {
	factory, err := GetFactory("mac")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()

	fmt.Println(button.Paint())
	fmt.Println(checkbox.Paint())
}
