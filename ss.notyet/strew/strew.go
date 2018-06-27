// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"time"

	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// ===========================================================================
// Beg of anyThingStrew - scatter them

// anyThingStrew returns a slice (of size = size) of channels
// one of which shall receive each inp before close.
func anyThingStrew(inp <-chan anyThing, size int) (outS [](<-chan anyThing)) {
	chaS := make([]chan anyThing, size)
	for i := 0; i < size; i++ {
		chaS[i] = make(chan anyThing)
	}

	go strewanyThing(inp, chaS...)

	outS = make([]<-chan anyThing, size)
	for i := 0; i < size; i++ {
		outS[i] = (<-chan anyThing)(chaS[i]) // convert `chan` to `<-chan`
	}

	return outS
}

// c strewanyThing(inp <-chan anyThing, outS ...chan<- anyThing) {
// Note: go does not convert the passed slice `[]chan anyThing` to `[]chan<- anyThing` automatically.
// So, we do neither here, as we are lazy (we just call an internal helper function).
func strewanyThing(inp <-chan anyThing, outS ...chan anyThing) {

	for i := range inp {
		for !trySendanyThing(i, outS...) {
			time.Sleep(time.Millisecond * 10) // wait a little before retry
		} // !sent
	} // inp

	for o := range outS {
		close(outS[o])
	}
}

func trySendanyThing(inp anyThing, outS ...chan anyThing) bool {

	for o := range outS {

		select { // try to send
		case outS[o] <- inp:
			return true
		default:
			// keep trying
		}

	} // outS
	return false
}

// End of anyThingStrew - scatter them
// ===========================================================================
