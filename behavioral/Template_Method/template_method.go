//Explanation:
//Template Structure (Beverage): Defines the PrepareRecipe method as a template method, incorporating steps like boiling water, brewing, pouring, and adding condiments.
//Concrete Implementations (Coffee, Tea): Provide specific implementations for the steps defined as part of the template method. They can customize parts of the procedure without impacting the overall algorithm.
//Main Program: Demonstrates how to prepare coffee and tea using the template method, showcasing how different steps are executed in a uniform sequence but with unique details.

package main

import "fmt"

// CaffeineBeverage defines the template method and its steps
type CaffeineBeverage interface {
	PrepareRecipe()
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
}

// Beverage is the common template structure
type Beverage struct {
	CaffeineBeverage
}

// PrepareRecipe defines the skeleton of a method
func (b *Beverage) PrepareRecipe() {
	b.BoilWater()
	b.Brew()
	b.PourInCup()
	b.AddCondiments()
}

// Coffee struct represents coffee preparation
type Coffee struct {
	Beverage
}

// NewCoffee initializes a new Coffee instance with the beverage template
func NewCoffee() *Coffee {
	c := &Coffee{}
	c.CaffeineBeverage = c // Struct embedding to use the methods
	return c
}

func (c *Coffee) BoilWater() {
	fmt.Println("Boiling water for coffee")
}

func (c *Coffee) Brew() {
	fmt.Println("Dripping Coffee through filter")
}

func (c *Coffee) PourInCup() {
	fmt.Println("Pouring coffee into cup")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Adding sugar and milk to coffee")
}

// Tea struct represents tea preparation
type Tea struct {
	Beverage
}

// NewTea initializes a new Tea instance with the beverage template
func NewTea() *Tea {
	t := &Tea{}
	t.CaffeineBeverage = t
	return t
}

func (t *Tea) BoilWater() {
	fmt.Println("Boiling water for tea")
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) PourInCup() {
	fmt.Println("Pouring tea into cup")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Adding lemon to tea")
}

func main() {
	coffee := NewCoffee()
	tea := NewTea()

	fmt.Println("Preparing Coffee:")
	coffee.PrepareRecipe()

	fmt.Println("\nPreparing Tea:")
	tea.PrepareRecipe()
}

//OUTPUT
//
//Preparing Coffee:
//Boiling water for coffee
//Dripping Coffee through filter
//Pouring coffee into cup
//Adding sugar and milk to coffee
//
//Preparing Tea:
//Boiling water for tea
//Steeping the tea
//Pouring tea into cup
//Adding lemon to tea
