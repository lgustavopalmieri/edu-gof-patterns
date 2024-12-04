//Explanation:
//Iterator Interface: Declares methods for iteration (HasNext and Next) that any iterator must implement.
//Concrete Iterator (NameIterator): Implements the Iterator interface and keeps track of the current iteration state using an index.
//Aggregate (NameCollection): Represents the collection that provides an Iterator method to create an iterator for itself.
//Client Code: Uses the iterator to traverse the collection of names, obtaining elements sequentially via the iterator's interface.

package main

import "fmt"

// Iterator interface defines methods for traversing a collection
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// NameCollection is a collection that holds a slice of names
type NameCollection struct {
	names []string
}

// NameIterator implements the Iterator interface for NameCollection
type NameIterator struct {
	collection *NameCollection
	index      int
}

// HasNext checks if there is a next element in the collection
func (n *NameIterator) HasNext() bool {
	return n.index < len(n.collection.names)
}

// Next returns the next element in the collection
func (n *NameIterator) Next() interface{} {
	if n.HasNext() {
		name := n.collection.names[n.index]
		n.index++
		return name
	}
	return nil
}

// Iterator returns a new iterator for the NameCollection
func (n *NameCollection) Iterator() Iterator {
	return &NameIterator{
		collection: n,
		index:      0,
	}
}

func main() {
	collection := &NameCollection{
		names: []string{"Alice", "Bob", "Charlie", "Diana"},
	}

	iterator := collection.Iterator()

	fmt.Println("Iterating over names:")
	for iterator.HasNext() {
		name := iterator.Next()
		fmt.Println(name)
	}
}

//OUTPUT
//
//Iterating over names:
//Alice
//Bob
//Charlie
//Diana
