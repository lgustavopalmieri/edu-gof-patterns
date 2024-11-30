//Explanation:
//Subject (Stock): Manages a list of observers and notifies them of any changes to the stock price.
//Observers (Investor): Implement the Observer interface and update their state in response to the subject's notifications.
//Notification: When the stock price changes, NotifyObservers is called to update all registered investors.
//Client Code: It registers investors (observers), changes stock prices, and observes notifications.

package main

import (
	"fmt"
)

// Observer defines an interface for objects that should be notified of changes
type Observer interface {
	Update(stockName string, price float64)
}

// Subject defines an interface to manage observers
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// Stock represents the subject
type Stock struct {
	observers map[Observer]struct{}
	name      string
	price     float64
}

// NewStock initializes a new Stock with a name
func NewStock(name string) *Stock {
	return &Stock{
		observers: make(map[Observer]struct{}),
		name:      name,
	}
}

// RegisterObserver adds an observer to the stock
func (s *Stock) RegisterObserver(observer Observer) {
	s.observers[observer] = struct{}{}
}

// RemoveObserver removes an observer from the stock
func (s *Stock) RemoveObserver(observer Observer) {
	delete(s.observers, observer)
}

// NotifyObservers notifies all observers about a state change
func (s *Stock) NotifyObservers() {
	for observer := range s.observers {
		observer.Update(s.name, s.price)
	}
}

// SetPrice changes the stock price and notifies observers
func (s *Stock) SetPrice(price float64) {
	s.price = price
	s.NotifyObservers()
}

// Investor represents a concrete observer
type Investor struct {
	name string
}

// Update is called when the subject changes
func (i *Investor) Update(stockName string, price float64) {
	fmt.Printf("Investor %s received update: %s is now $%.2f\n", i.name, stockName, price)
}

func main() {
	appleStock := NewStock("Apple")
	investor1 := &Investor{name: "John"}
	investor2 := &Investor{name: "Jane"}

	appleStock.RegisterObserver(investor1)
	appleStock.RegisterObserver(investor2)

	// Change the stock price and notify observers
	appleStock.SetPrice(150.0)
	appleStock.SetPrice(155.0)

	// Remove one observer and change the price
	appleStock.RemoveObserver(investor1)
	appleStock.SetPrice(160.0)
}
