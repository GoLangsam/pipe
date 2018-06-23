// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// anyOwner is the generic who shall own the methods.
//  Note: Need to use `generic.Number` here as `generic.Type` is an interface and cannot have any method.
type anyOwner generic.Number

// ===========================================================================
// Beg of anyThingSame comparator

// inspired by go/doc/play/tree.go

// anyThingSame reads values from two channels in lockstep
// and iff they have the same contents then
// `true` is sent on the returned bool channel
// before close.
func (my anyOwner) anyThingSame(same func(a, b anyThing) bool, inp1, inp2 <-chan anyThing) (out <-chan bool) {
	cha := make(chan bool)
	go my.sameanyThing(cha, same, inp1, inp2)
	return cha
}

func (my anyOwner) sameanyThing(out chan<- bool, same func(a, b anyThing) bool, inp1, inp2 <-chan anyThing) {
	defer close(out)
	for {
		v1, ok1 := <-inp1
		v2, ok2 := <-inp2

		if !ok1 || !ok2 {
			out <- ok1 == ok2
			return
		}
		if !same(v1, v2) {
			out <- false
			return
		}
	}
}

// End of anyThingSame comparator
// ===========================================================================
