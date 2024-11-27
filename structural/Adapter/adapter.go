//Explanation:
//Target Interface: Represents the interface that the client expects.
//Adaptee: A class or service with an incompatible interface that needs to be adapted.
//Adapter: Implements the Target interface and internally holds an instance of Adaptee. It translates calls to the Target interface into calls to the Adaptee's SpecificRequest method.
//Client Code: Uses the Target interface, interacting with the Adaptee indirectly through the Adapter.

package main

import "fmt"

// Target is the interface that your application expects
type Target interface {
	Request() string
}

// Adaptee is the service with a different interface
type Adaptee struct{}

// SpecificRequest returns data in a format that doesn't match the Target interface
func (a *Adaptee) SpecificRequest() string {
	return "Adaptee's specific request"
}

// Adapter makes the Adaptee's interface compatible with the Target interface
type Adapter struct {
	adaptee *Adaptee
}

// Request adapts the SpecificRequest method of Adaptee to fit the Target interface
func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	// Create an instance of Adaptee
	adaptee := &Adaptee{}

	// Wrap it with an Adapter
	adapter := &Adapter{adaptee: adaptee}

	// Use the adapter through the Target interface
	fmt.Println(adapter.Request())
}
