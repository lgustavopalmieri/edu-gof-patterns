//Explanation:
//Strategy Interface (SortingStrategy): Defines a method Sort that all concrete strategies must implement.
//Concrete Strategies (BubbleSort, QuickSort): Provide specific implementations of the Sort method.
//Context (Context): Holds a reference to a SortingStrategy and is able to change the strategy at runtime.
//Client Code: Chooses the strategy and asks the context to carry out the algorithm using the selected strategy.

package main

import (
	"fmt"
	"sort"
)

// SortingStrategy is the strategy interface
type SortingStrategy interface {
	Sort([]int)
}

// BubbleSort is a concrete strategy implementing bubble sort
type BubbleSort struct{}

func (b *BubbleSort) Sort(data []int) {
	n := len(data)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	fmt.Println("Sorted using Bubble Sort:", data)
}

// QuickSort is a concrete strategy using the standard Go library's sort
type QuickSort struct{}

func (q *QuickSort) Sort(data []int) {
	sort.Ints(data)
	fmt.Println("Sorted using Quick Sort:", data)
}

// Context maintains a reference to a strategy object
type Context struct {
	strategy SortingStrategy
}

// SetStrategy sets the current strategy
func (c *Context) SetStrategy(strategy SortingStrategy) {
	c.strategy = strategy
}

// ExecuteSort performs sorting using the current strategy
func (c *Context) ExecuteSort(data []int) {
	c.strategy.Sort(data)
}

func main() {
	data := []int{33, 10, 55, 71, 38, 19}

	bubbleSort := &BubbleSort{}
	quickSort := &QuickSort{}

	context := &Context{}

	// Use BubbleSort strategy
	context.SetStrategy(bubbleSort)
	context.ExecuteSort(append([]int(nil), data...)) // Work on a copy of data

	// Use QuickSort strategy
	context.SetStrategy(quickSort)
	context.ExecuteSort(append([]int(nil), data...)) // Work on a copy of data
}
