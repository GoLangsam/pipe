// from go/doc/play/tree.go

package pipe

import (
	"fmt"
)

// Go's concurrency primitives make it easy to
// express concurrent concepts, such as
// this binary tree comparison.
//
// Trees may be of different shapes,
// but have the same contents. For example:
//
//        4               6
//      2   6          4     7
//     1 3 5 7       2   5
//                  1 3
//
// This program compares a pair of trees by
// walking each in its own goroutine,
// sending their contents through a channel
// to a third goroutine that compares them.
func ExampleanyThingFrom_Same_tree() {

	// same reports iff a and b are equal
	same := func(a, b anyThing) bool {
		v1 := a.(int)
		v2 := b.(int)
		return v1 == v2
	}

	// Compare reads values from two Walkers
	// that run simultaneously, and returns true
	// if t1 and t2 have the same contents.
	Compare := func(t1, t2 *Tree) bool {
		c1, c2 := Walker(t1), Walker(t2)
		return <-c1.Same(same, c2)
	}

	t1 := New(100, 1)
	fmt.Println(Compare(t1, New(100, 1)), "Same Contents")
	fmt.Println(Compare(t1, New(99, 1)), "Differing Sizes")
	fmt.Println(Compare(t1, New(100, 2)), "Differing Values")
	fmt.Println(Compare(t1, New(101, 2)), "Dissimilar")
	// Output:
	// true Same Contents
	// false Differing Sizes
	// false Differing Values
	// false Dissimilar
}
