package main

import "fmt"

func main() {
	a := [4]int{1, -2, -3, 0}

	// The slice `s` is a view of the array `a`.
	s := a[:3] // s: [1 -2 -3], a: [1 -2 -3 0]
	printSlices(s, a[:])

	// The slice `s` points to the same memory location as the array `a`.
	// Modifying the array `a` reflects to its slice `s`.
	a[1] = 2 // s: [1 2 -3], a: [1 2 -3 0]
	printSlices(s, a[:])

	// Modifying the slice `s` reflects to the backing array.
	s[2] = 3 // s: [1 2 3 4], a: [1 2 3 4]
	printSlices(s, a[:])

	// The slice can be appended! As long as the capacity is enough, the slice
	// `s` will still reference the same array `a`, hence the backing array `a`
	// is 'affected'.
	s = append(s, 4) // s: [1 2 3 4], a: [1 2 3 4]
	printSlices(s, a[:])

	// Since the capacity is full, a new array is created to back the slice `s`,
	// 'abandoning' its backing array `a`.
	s = append(s, 5) // s: [1 2 3 4 5], a: [1 2 3 4]
	printSlices(s, a[:])

	// Now, `a` is no more in sync with `s`.
	a[0] = 10 // s: [1 2 3 4 5], a: [10 2 3 4]
	printSlices(s, a[:])

	// The question is: why did they name it 'slice' if it can be appended to,
	// modifying its backing array, and when it runs out of capacity, it can
	// discard the original backing array?
	//
	// The answer may be: it is more similar to the "vector" in C++, which is
	// an array-like data structure that allows efficient resizing and
	// appending.
	//
	// However, if that were the sole purpose, they shouldn't allow us to
	// explicitly specify which array backs the slice, as this breaks the
	// abstraction level.
}

// Print the slices `s` and `a`, including their length, capacity, and content.
func printSlices(s []int, a []int) {
	fmt.Printf("s: len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("a: len=%d cap=%d %v\n", len(a), cap(a), a)
	fmt.Println()
}
