//Explanation:
//State Interface (State): Declares methods for handling state-specific behavior, implemented by each concrete state.
//Concrete States (OnState, OffState): Implement the state-specific behavior and transition logic.
//Context (Device): Maintains an instance of the current state and delegates state-specific requests to this object. Allows state transitions by changing its current state.
//Client Code: Interacts with the Device, triggering state changes through the OnButtonPressed and OffButtonPressed methods, which delegate behavior to the current state.

package main

import "fmt"

// State is the interface for device states
type State interface {
	OnButtonPressed(*Device)
	OffButtonPressed(*Device)
}

// Device holds current state and provides the context for state transitions
type Device struct {
	currentState State
}

// NewDevice initializes the device with an off state
func NewDevice() *Device {
	return &Device{currentState: &OffState{}}
}

// SetState changes the device's current state
func (d *Device) SetState(state State) {
	d.currentState = state
}

// OnButtonPressed triggers behavior based on the current state
func (d *Device) OnButtonPressed() {
	d.currentState.OnButtonPressed(d)
}

// OffButtonPressed triggers behavior based on the current state
func (d *Device) OffButtonPressed() {
	d.currentState.OffButtonPressed(d)
}

// OnState represents a state where the device is on
type OnState struct{}

func (s *OnState) OnButtonPressed(d *Device) {
	fmt.Println("Device is already ON.")
}

func (s *OnState) OffButtonPressed(d *Device) {
	fmt.Println("Turning device OFF.")
	d.SetState(&OffState{})
}

// OffState represents a state where the device is off
type OffState struct{}

func (s *OffState) OnButtonPressed(d *Device) {
	fmt.Println("Turning device ON.")
	d.SetState(&OnState{})
}

func (s *OffState) OffButtonPressed(d *Device) {
	fmt.Println("Device is already OFF.")
}

func main() {
	device := NewDevice()

	// Triggering state changes
	device.OnButtonPressed()  // Turns ON
	device.OnButtonPressed()  // Already ON
	device.OffButtonPressed() // Turns OFF
	device.OffButtonPressed() // Already OFF
	device.OnButtonPressed()  // Turns ON
}
