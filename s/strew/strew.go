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
	chaS := make(map[chan anyThing]struct{}, size)
	for i := 0; i < size; i++ {
		chaS[make(chan anyThing)] = struct{}{}
	}

	go strewanyThing(inp, chaS)

	outS = make([]<-chan anyThing, size)
	i := 0
	for c := range chaS {
		outS[i] = (<-chan anyThing)(c) // convert `chan` to `<-chan`
		i++
	}

	return outS
}

// c strewanyThing(inp <-chan anyThing, outS ...chan<- anyThing) {
// Note: go does not convert the passed slice `[]chan anyThing` to `[]chan<- anyThing` automatically.
// So, we do neither here, as we are lazy (we just call an internal helper function).
func strewanyThing(inp <-chan anyThing, outS map[chan anyThing]struct{}) {

	for i := range inp {
		for !trySendanyThing(i, outS) {
			time.Sleep(time.Millisecond * 10) // wait a little before retry
		} // !sent
	} // inp

	for o := range outS {
		close(o)
	}
}

func trySendanyThing(inp anyThing, outS map[chan anyThing]struct{}) bool {

	for o := range outS {

		select { // try to send
		case o <- inp:
			return true
		default:
			// keep trying
		}

	} // outS
	return false
}

// End of anyThingStrew - scatter them
// ===========================================================================
