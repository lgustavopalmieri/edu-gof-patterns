//Explanation:
//Product (House): The complex object being constructed, which has various parts like walls, roof, and floor.
//Builder Interface (HouseBuilder): Defines the steps required to build different parts of the House.
//Concrete Builder (ConcreteBuilder): Implements the HouseBuilder interface and constructs the actual parts of the House. Stores the complex object being built (house).
//Director: Orchestrates the construction process using a HouseBuilder. It knows the order in which to execute the building steps.
//Client Code: Uses the Director and a specific ConcreteBuilder to construct a House and then retrieves the final product.

package main

import "fmt"

// House represents the complex object
type House struct {
	Walls string
	Roof  string
	Floor string
}

// HouseBuilder defines the steps to build a house
type HouseBuilder interface {
	BuildWalls()
	BuildRoof()
	BuildFloor()
	GetHouse() House
}

// ConcreteBuilder implements the HouseBuilder interface
type ConcreteBuilder struct {
	house House
}

// BuildWalls constructs walls for the house
func (b *ConcreteBuilder) BuildWalls() {
	b.house.Walls = "Wooden Walls"
}

// BuildRoof constructs a roof for the house
func (b *ConcreteBuilder) BuildRoof() {
	b.house.Roof = "Concrete Roof"
}

// BuildFloor constructs a floor for the house
func (b *ConcreteBuilder) BuildFloor() {
	b.house.Floor = "Marble Floor"
}

// GetHouse returns the constructed house
func (b *ConcreteBuilder) GetHouse() House {
	return b.house
}

// Director constructs a house using HouseBuilder
type Director struct {
	builder HouseBuilder
}

// NewDirector initializes a Director with a builder
func NewDirector(b HouseBuilder) *Director {
	return &Director{
		builder: b,
	}
}

// Construct builds a house using the builder
func (d *Director) Construct() {
	d.builder.BuildWalls()
	d.builder.BuildRoof()
	d.builder.BuildFloor()
}

func main() {
	// Initialize a concrete builder
	builder := &ConcreteBuilder{}

	// Create a director with the builder
	director := NewDirector(builder)

	// Construct the house
	director.Construct()

	// Retrieve the built house
	house := builder.GetHouse()

	// Output the details of the house
	fmt.Printf("House with %s, %s, and %s.\n", house.Walls, house.Roof, house.Floor)
}
