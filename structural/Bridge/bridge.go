//Explanation:
//Implementation Interface (MessageSender): Defines the method SendMessage that any communication method (e.g., SMS, Email) must implement.
//Concrete Implementations (SMS, Email): Provide specific means of sending messages by implementing the MessageSender interface.
//Abstraction Interface (Notification): Represents the high-level control of the notification process.
//Concrete Abstractions (AlertNotification, RemindNotification): Extend the abstraction and use a MessageSender to perform the actual task, demonstrating the decoupling of abstraction from implementation.
//Client Code: Configures the bridge structure in main by using different implementations for different abstractions, providing flexibility and supporting runtime changes.

package main

import "fmt"

// MessageSender is the implementation interface
type MessageSender interface {
	SendMessage(msg string) string
}

// SMS is a concrete implementation of the MessageSender interface
type SMS struct{}

// SendMessage implementation for SMS
func (s *SMS) SendMessage(msg string) string {
	return fmt.Sprintf("Sending SMS with message: %s", msg)
}

// Email is another concrete implementation of the MessageSender interface
type Email struct{}

// SendMessage implementation for Email
func (e *Email) SendMessage(msg string) string {
	return fmt.Sprintf("Sending Email with message: %s", msg)
}

// Notification is the abstraction interface
type Notification interface {
	Notify(message string) string
}

// AlertNotification uses a MessageSender to send notifications
type AlertNotification struct {
	sender MessageSender
}

// Notify sends an alert message using the MessageSender
func (n *AlertNotification) Notify(message string) string {
	return n.sender.SendMessage(message)
}

// RemindNotification uses a MessageSender to send reminders
type RemindNotification struct {
	sender MessageSender
}

// Notify sends a reminder message using the MessageSender
func (r *RemindNotification) Notify(message string) string {
	return r.sender.SendMessage(message)
}

func main() {
	sms := &SMS{}
	email := &Email{}

	alert := &AlertNotification{sender: sms}
	reminder := &RemindNotification{sender: email}

	fmt.Println(alert.Notify("This is an alert"))
	fmt.Println(reminder.Notify("This is a reminder"))
}
