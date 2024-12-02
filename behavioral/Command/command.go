//Explanation:
//Command Interface (Command): Defines Execute and Undo methods for any command.
//Receiver Classes (Television, AirConditioner): Perform operations-specific tasks on the devices.
//Concrete Commands:
//TVOnCommand and TVOffCommand encapsulate operations to turn the TV on and off.
//ACSetTemperatureCommand encapsulates setting the AC's temperature and remembers previous states for undo functionality.
//Invoker (RemoteControl): Manages command execution and maintains a history to undo operations.
//Client Code: Controls the TV and AC, showing command execution and reversal.

package main

import "fmt"

// Command interface defines the Execute and Undo methods
type Command interface {
	Execute()
	Undo()
}

// Television is a receiver class that represents a TV
type Television struct {
	isOn bool
}

func (tv *Television) On() {
	tv.isOn = true
	fmt.Println("Television is on")
}

func (tv *Television) Off() {
	tv.isOn = false
	fmt.Println("Television is off")
}

// AirConditioner is a receiver class for air conditioner operations
type AirConditioner struct {
	temperature int
}

func (ac *AirConditioner) SetTemperature(temp int) {
	ac.temperature = temp
	fmt.Printf("Air conditioner temperature set to %d°C\n", ac.temperature)
}

// TVOnCommand is a concrete command to turn the TV on
type TVOnCommand struct {
	tv *Television
}

func (c *TVOnCommand) Execute() {
	c.tv.On()
}

func (c *TVOnCommand) Undo() {
	c.tv.Off()
}

// TVOffCommand is a concrete command to turn the TV off
type TVOffCommand struct {
	tv *Television
}

func (c *TVOffCommand) Execute() {
	c.tv.Off()
}

func (c *TVOffCommand) Undo() {
	c.tv.On()
}

// ACSetTemperatureCommand sets the temperature of the air conditioner
type ACSetTemperatureCommand struct {
	ac       *AirConditioner
	prevTemp int
	newTemp  int
}

func (c *ACSetTemperatureCommand) Execute() {
	c.prevTemp = c.ac.temperature
	c.ac.SetTemperature(c.newTemp)
}

func (c *ACSetTemperatureCommand) Undo() {
	c.ac.SetTemperature(c.prevTemp)
}

// RemoteControl is the invoker
type RemoteControl struct {
	history []Command
}

func (r *RemoteControl) PressButton(command Command) {
	command.Execute()
	r.history = append(r.history, command)
}

func (r *RemoteControl) PressUndo() {
	if len(r.history) == 0 {
		fmt.Println("No commands to undo")
		return
	}
	lastCommand := r.history[len(r.history)-1]
	lastCommand.Undo()
	r.history = r.history[:len(r.history)-1]
}

func main() {
	tv := &Television{}
	ac := &AirConditioner{temperature: 24} // Default temperature is 24°C

	tvOn := &TVOnCommand{tv: tv}
	tvOff := &TVOffCommand{tv: tv}
	acTemp := &ACSetTemperatureCommand{ac: ac, newTemp: 18}

	remote := &RemoteControl{}

	// Control the TV
	remote.PressButton(tvOn)
	remote.PressButton(tvOff)
	remote.PressUndo()

	// Control the Air Conditioner
	remote.PressButton(acTemp)
	remote.PressUndo()
}

//OUTPUT:
//
//Television is on
//Television is off
//Television is on
//Air conditioner temperature set to 18°C
//Air conditioner temperature set to 24°C
