// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of Fork functions

// Fork returns two channels
// either of which is to receive
// every result of inp
// before close.
func (inp anyThingFrom) Fork() (out1, out2 anyThingFrom) {
	cha1 := make(chan anyThing)
	cha2 := make(chan anyThing)
	go inp.fork(cha1, cha2)
	return cha1, cha2
}

/* not used - kept for reference only.
func (inp anyThingFrom) fork(out1, out2 anyThingInto) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
} */

func (inp anyThingFrom) fork(out1, out2 anyThingInto) {
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

// End of Fork functions
// ===========================================================================
