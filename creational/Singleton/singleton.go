package main

import (
	"fmt"
	"sync"
)

// Singleton is the type that will be a singleton
type Singleton struct {
	field1 byte

	// Add fields as you need
}

// instance is the single instance of Singleton
var instance *Singleton
var once sync.Once

// getInstance provides global access to the Singleton instance
func getInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	// Retrieve the singleton instance
	s1 := getInstance()
	s2 := getInstance()

	if s1 == s2 {
		fmt.Println("Same instance")
	} else {
		fmt.Println("Different instances")
	}
}
