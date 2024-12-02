//Explanation:
//Mediator Interface (Mediator): Defines a method for sending messages.
//Concrete Mediator (ChatRoom): Implements the mediator interface to coordinate communication between users by managing user lists and dispatching messages.
//Colleague (User): Represents the objects that send and receive messages. Users are aware only of the mediator and not each other.
//Client Code: Sets up the users and chat room (mediator), demonstrating how communication between users is facilitated by the mediator.

package main

import "fmt"

// Mediator interface defines the communication protocol
type Mediator interface {
	SendMessage(msg string, user *User)
}

// ChatRoom is the concrete mediator class
type ChatRoom struct {
	users []*User
}

// AddUser adds a user to the chat room
func (c *ChatRoom) AddUser(user *User) {
	c.users = append(c.users, user)
}

// SendMessage sends a message to all users in the chat room except the sender
func (c *ChatRoom) SendMessage(msg string, user *User) {
	for _, u := range c.users {
		if u != user {
			u.Receive(msg)
		}
	}
}

// User represents a participant in the chat room
type User struct {
	name     string
	mediator Mediator
}

// NewUser creates a new chat room user
func NewUser(name string, mediator Mediator) *User {
	return &User{name: name, mediator: mediator}
}

// Send sends a message via the mediator
func (u *User) Send(msg string) {
	fmt.Printf("[%s]: Sending message: %s\n", u.name, msg)
	u.mediator.SendMessage(msg, u)
}

// Receive receives a message
func (u *User) Receive(msg string) {
	fmt.Printf("[%s]: Received message: %s\n", u.name, msg)
}

func main() {
	chatroom := &ChatRoom{}

	user1 := NewUser("Alice", chatroom)
	user2 := NewUser("Bob", chatroom)
	user3 := NewUser("Charlie", chatroom)

	chatroom.AddUser(user1)
	chatroom.AddUser(user2)
	chatroom.AddUser(user3)

	user1.Send("Hello, everyone!")
	user2.Send("Hi Alice!")
}

//OUTPUT:
//
//[Alice]: Sending message: Hello, everyone!
//[Bob]: Received message: Hello, everyone!
//[Charlie]: Received message: Hello, everyone!
//[Bob]: Sending message: Hi Alice!
//[Alice]: Received message: Hi Alice!
//[Charlie]: Received message: Hi Alice!
