// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// anyThingFrom is a receive-only anyThing channel
type anyThingFrom <-chan anyThing

// anyThingInto is a send-only anyThing channel
type anyThingInto chan<- anyThing

// ===========================================================================
// Beg of anyThingSame comparator

// inspired by go/doc/play/tree.go

// anyThingSame reads values from two channels in lockstep
// and iff they have the same contents then
// `true` is sent on the returned bool channel
// before close.
func (inp anyThingFrom) anyThingSame(same func(a, b anyThing) bool, inp2 anyThingFrom) (out <-chan bool) {
	cha := make(chan bool)
	go inp.sameanyThing(cha, same, inp2)
	return cha
}

func (inp anyThingFrom) sameanyThing(out chan<- bool, same func(a, b anyThing) bool, inp2 anyThingFrom) {
	defer close(out)
	for {
		v1, ok1 := <-inp
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
