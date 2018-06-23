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
// Beg of anyThingFanOut

// anyThingFanOut returns a slice (of size = size) of channels
// each of which shall receive any inp before close.
func (my anyOwner) anyThingFanOut(inp <-chan anyThing, size int) (outS [](<-chan anyThing)) {
	chaS := make([]chan anyThing, size)
	for i := 0; i < size; i++ {
		chaS[i] = make(chan anyThing)
	}

	go my.fananyThingOut(inp, chaS...)

	outS = make([]<-chan anyThing, size)
	for i := 0; i < size; i++ {
		outS[i] = chaS[i] // convert `chan` to `<-chan`
	}

	return outS
}

// c (my anyOwner) fananyThingOut(inp <-chan anyThing, outs ...chan<- anyThing) {
func (my anyOwner) fananyThingOut(inp <-chan anyThing, outs ...chan anyThing) {

	for i := range inp {
		for o := range outs {
			outs[o] <- i
		}
	}

	for o := range outs {
		close(outs[o])
	}

}

// End of anyThingFanOut
// ===========================================================================
