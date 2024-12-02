//Explanation:
//SupportHandler Interface: Defines methods for setting the next handler and handling requests.
//BaseHandler: Provides common functionality for handling the next handler in the chain.
//Concrete Handlers (LevelOneSupport, LevelTwoSupport, LevelThreeSupport): Implement specific handling logic, either processing the request or passing it to the next handler.
//Client Code: Sends requests to the first handler in the chain and relies on it to pass the request through the chain until it is handled.

package main

import "fmt"

// SupportHandler defines the interface for handling requests
type SupportHandler interface {
	SetNext(handler SupportHandler)
	HandleRequest(issueLevel int)
}

// BaseHandler defines a base handler with common functionality
type BaseHandler struct {
	nextHandler SupportHandler
}

func (h *BaseHandler) SetNext(handler SupportHandler) {
	h.nextHandler = handler
}

func (h *BaseHandler) HandleRequest(issueLevel int) {
	if h.nextHandler != nil {
		h.nextHandler.HandleRequest(issueLevel)
	}
}

// LevelOneSupport handles Level 1 issues
type LevelOneSupport struct {
	BaseHandler
}

func (h *LevelOneSupport) HandleRequest(issueLevel int) {
	if issueLevel == 1 {
		fmt.Println("Level 1 Support: Handling issue.")
	} else {
		fmt.Println("Level 1 Support: Passing issue to next level.")
		h.BaseHandler.HandleRequest(issueLevel)
	}
}

// LevelTwoSupport handles Level 2 issues
type LevelTwoSupport struct {
	BaseHandler
}

func (h *LevelTwoSupport) HandleRequest(issueLevel int) {
	if issueLevel == 2 {
		fmt.Println("Level 2 Support: Handling issue.")
	} else {
		fmt.Println("Level 2 Support: Passing issue to next level.")
		h.BaseHandler.HandleRequest(issueLevel)
	}
}

// LevelThreeSupport handles Level 3 issues
type LevelThreeSupport struct {
	BaseHandler
}

func (h *LevelThreeSupport) HandleRequest(issueLevel int) {
	if issueLevel == 3 {
		fmt.Println("Level 3 Support: Handling issue.")
	} else {
		fmt.Println("Level 3 Support: Unable to handle issue.")
	}
}

func main() {
	// Create handlers
	level1 := &LevelOneSupport{}
	level2 := &LevelTwoSupport{}
	level3 := &LevelThreeSupport{}

	// Set up the chain
	level1.SetNext(level2)
	level2.SetNext(level3)

	// Client sends requests
	fmt.Println("Sending Level 1 issue:")
	level1.HandleRequest(1)

	fmt.Println("\nSending Level 2 issue:")
	level1.HandleRequest(2)

	fmt.Println("\nSending Level 3 issue:")
	level1.HandleRequest(3)

	fmt.Println("\nSending unknown issue level:")
	level1.HandleRequest(4)
}

//OUTPUT:
//
//Sending Level 1 issue:
//Level 1 Support: Handling issue.
//
//Sending Level 2 issue:
//Level 1 Support: Passing issue to next level.
//Level 2 Support: Handling issue.
//
//Sending Level 3 issue:
//Level 1 Support: Passing issue to next level.
//Level 2 Support: Passing issue to next level.
//Level 3 Support: Handling issue.
//
//Sending unknown issue level:
//Level 1 Support: Passing issue to next level.
//Level 2 Support: Passing issue to next level.
//Level 3 Support: Unable to handle issue.
