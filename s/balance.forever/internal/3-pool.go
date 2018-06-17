// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package balance

// ===========================================================================
// Beg of Pool

// Pool is a slice of (pointers to) Worker and
// implements Heap: Len Less Swap Push Pop.
type Pool []*Worker

// Len reports the number of elements in the heap.
func (p *Pool) Len() int { return len(*p) }

// Less reports whether the element @ [i] should sort before the element @ [j].
func (p *Pool) Less(i, j int) bool { return (*p)[i].pending < (*p)[j].pending }

// Swap swaps the elements @ [i] and [j].
func (p *Pool) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i] // swap Elements
	(*p)[i].index, (*p)[j].index = i, j // adjust indices
}

// Push add v as element @ [ Len() ].
func (p *Pool) Push(v interface{}) {
	w := v.(*Worker)
	w.index = len(*p)
	*p = append(*p, w)
}

// Pop removes and returns the element @ [ Len() - 1 ].
func (p *Pool) Pop() (v interface{}) {
	*p, v = (*p)[:p.Len()-1], (*p)[p.Len()-1]
	return
}

// End of Pool
// ===========================================================================
