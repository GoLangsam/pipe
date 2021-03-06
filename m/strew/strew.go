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
// Beg of Strew - scatter them

// Strew returns a slice (of size = size) of channels
// one of which shall receive each inp before close.
func (inp anyThingFrom) Strew(size int) (outS []anyThingFrom) {
	chaS := make(map[chan anyThing]struct{}, size)
	for i := 0; i < size; i++ {
		chaS[make(chan anyThing)] = struct{}{}
	}

	go inp.strew(chaS)

	outS = make([]anyThingFrom, size)
	i := 0
	for c := range chaS {
		outS[i] = (anyThingFrom)(c) // convert `chan anyThing` to anyThingFrom
		i++
	}

	return outS
}

func (inp anyThingFrom) strew(outS map[chan anyThing]struct{}) {

	for i := range inp {
		for !inp.trySend(i, outS) {
			time.Sleep(time.Millisecond * 10) // wait a little before retry
		} // !sent
	} // inp

	for o := range outS {
		close(o)
	}
}

func (static anyThingFrom) trySend(inp anyThing, outS map[chan anyThing]struct{}) bool {

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

// End of Strew - scatter them
// ===========================================================================
