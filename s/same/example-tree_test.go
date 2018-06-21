// from go/doc/play/tree.go

package pipe

import (
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk traverses a tree depth-first,
// sending each Value on a channel.
func Walk(t *Tree, ch chan anyThing) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func Walker(t *Tree) <-chan anyThing {
	ch := make(chan anyThing)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

// New returns a new, random binary tree
// holding the values 1k, 2k, ..., nk.
func New(n, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}
