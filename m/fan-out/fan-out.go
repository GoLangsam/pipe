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
// Beg of anyThingFanOut

// anyThingFanOut returns a slice (of size = size) of channels
// each of which shall receive any inp before close.
func (inp anyThingFrom) anyThingFanOut(size int) (outS [](anyThingFrom)) {
	chaS := make([]chan anyThing, size)
	for i := 0; i < size; i++ {
		chaS[i] = make(chan anyThing)
	}

	go inp.fananyThingOut(chaS...)

	outS = make([]anyThingFrom, size)
	for i := 0; i < size; i++ {
		outS[i] = (anyThingFrom)(chaS[i]) // convert `chan` to `<-chan`
	}

	return outS
}

// c (inp anyThingFrom) fananyThingOut(outs ...anyThingInto) {
func (inp anyThingFrom) fananyThingOut(outs ...chan anyThing) {

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
