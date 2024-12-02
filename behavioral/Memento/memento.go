//Explanation:
//Memento Struct: Holds the state of the Editor at a particular instant. It is created and used by the editor to store and retrieve its state.
//Originator (Editor): The object that creates a memento containing its current state. It can also restore its state from the memento.
//Caretaker: Manages the collection of Memento objects, adding new mementos to history and providing them for state restoration when needed.
//Client Code: Simulates a user writing to the Editor, saving states, and restoring earlier states to implement undo functionality.

package main

import (
	"fmt"
)

// Memento stores the state of the Editor
type Memento struct {
	state string
}

// Editor is the originator, which can create and restore mementos
type Editor struct {
	content string
}

func (e *Editor) Write(text string) {
	e.content += text
}

func (e *Editor) Save() *Memento {
	return &Memento{state: e.content}
}

func (e *Editor) Restore(m *Memento) {
	e.content = m.state
}

func (e *Editor) Content() string {
	return e.content
}

// Caretaker manages the memento objects
type Caretaker struct {
	history []*Memento
}

func (c *Caretaker) AddMemento(m *Memento) {
	c.history = append(c.history, m)
}

func (c *Caretaker) GetMemento(index int) *Memento {
	if index >= 0 && index < len(c.history) {
		return c.history[index]
	}
	return nil
}

func main() {
	editor := &Editor{}
	caretaker := &Caretaker{}

	// Initial writing
	editor.Write("First line.\n")
	caretaker.AddMemento(editor.Save())

	// Additional writing
	editor.Write("Second line.\n")
	caretaker.AddMemento(editor.Save())

	// Writing further
	editor.Write("Third line.\n")
	fmt.Println("Current Content:\n", editor.Content())

	// Undoing operations
	editor.Restore(caretaker.GetMemento(1))
	fmt.Println("\nAfter Undo 1:\n", editor.Content())

	editor.Restore(caretaker.GetMemento(0))
	fmt.Println("\nAfter Undo 2:\n", editor.Content())
}

//OUTPUT:
//
//Current Content:
//First line.
//Second line.
//Third line.
//
//After Undo 1:
//First line.
//Second line.
//
//After Undo 2:
//First line.
