//Explanation:
//Vehicle Interface: Defines a Drive method that all vehicle types will implement.
//Concrete Types (Car, Bike): Implement the Vehicle interface by providing their own Drive method.
//VehicleFactory: Serves as the factory method that determines which type of Vehicle to instantiate based on the input parameter. This demonstrates polymorphic behavior by returning a specific vehicle type.
//Usage in main: By calling VehicleFactory, you can obtain a Vehicle instance at runtime based on specified criteria without knowing the specific class of object that will be returned.
//This approach maintains loose coupling by relying on the Vehicle interface rather than any concrete type, allowing the code to remain flexible and adaptable to future changes or additions of new vehicle types.

package main

import "fmt"

// Vehicle is the interface that all products must implement
type Vehicle interface {
	Drive() string
}

// Car struct implements the Vehicle interface
type Car struct{}

// Drive implements the Drive method of Vehicle interface for Car
func (c *Car) Drive() string {
	return "Driving a car"
}

// Bike struct implements the Vehicle interface
type Bike struct{}

// Drive implements the Drive method of Vehicle interface for Bike
func (b *Bike) Drive() string {
	return "Riding a bike"
}

// vehicleFactory function that acts as a factory method
func vehicleFactory(vehicleType string) (Vehicle, error) {
	switch vehicleType {
	case "car":
		return &Car{}, nil
	case "bike":
		return &Bike{}, nil
	default:
		return nil, fmt.Errorf("unknown vehicle type")
	}
}

func main() {
	vehicle, err := vehicleFactory("car")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(vehicle.Drive())

	vehicle, err = vehicleFactory("bike")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(vehicle.Drive())
}
