// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingPair functions

// anyThingPair returns a pair of channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func (inp anyThingFrom) anyThingPair() (out1, out2 anyThingFrom) {
	cha1 := make(chan anyThing)
	cha2 := make(chan anyThing)
	go inp.pairanyThing(cha1, cha2)
	return cha1, cha2
}

/* not used - kept for reference only.
func (inp anyThingFrom) pairanyThing(out1, out2 anyThingInto, inp anyThingFrom) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func (inp anyThingFrom) pairanyThing(out1, out2 anyThingInto) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		select { // send first to whomever is ready to receive
		case out1 <- i:
			out2 <- i
		case out2 <- i:
			out1 <- i
		}
	}
}

// End of anyThingPair functions
// ===========================================================================
