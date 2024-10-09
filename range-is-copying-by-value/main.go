package main

import "fmt"

func main() {
	// ========================================================================
	// DISCLAIMER: This is an opinionated criticism and may not reflect general
	// views. I could only find a few similar opinions while searching. Others
	// have found that the issue I am addressing here might not be significant
	// enough to affect performance in most cases.
	// ========================================================================

	// The implementation of `range` in Go might seem inefficient or not very
	// straightforward to use in an efficient way. Let's take a look:

	// Traverse over the slice using `range`
	s := []int{0, 1, 2, 3, 4}
	for i, v := range s {
		// Output: (0, 0), (1, 1), ...
		fmt.Printf("(%d, %d)\n", i, v)
	}

	type Foo struct{ foo int }

	// Initialize a slice of structs
	fooSlice := make([]Foo, 5)
	for i := 0; i < 5; i++ {
		fooSlice[i].foo = i
	}

	// Traverse using `range`
	for i, v := range fooSlice {
		// Output: (0, {0}), (1, {1}), ...
		fmt.Printf("(%d, %d)\n", i, v)
	}

	// Attempt to modify each entry using `range`
	for _, v := range fooSlice {
		v.foo = -1 // will this modify the original value?
	}

	// Since `range` performs 'copy by value', the original slice is not
	// affected.
	for i, v := range fooSlice {
		// Not really: (0, {-1}), (1, {-1}), ...
		// In fact: (0, {0}), (1, {1}), ...
		fmt.Printf("(%d, %d)\n", i, v)
	}

	// Instead, you can traverse like this:
	for i /*, v*/ := range fooSlice { // the value can be omitted
		fooSlice[i].foo = -1
	}
	for i, v := range fooSlice {
		// Output: (0, {-1}), (1, {-1}), ...
		fmt.Printf("(%d, %d)\n", i, v)
	}

	// Therefore, if the struct payload is large, using `range` for traversal
	// can be inefficient, unless each entry is accessed directly by its index,
	// which may be less straightforward compared to other languages that also
	// support range-based for loops natively.
}
