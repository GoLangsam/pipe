// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingFork functions

// anyThingFork returns two channels
// either of which is to receive
// every result of inp
// before close.
func anyThingFork(inp chan anyThing) (chan anyThing, chan anyThing) {
	out1 := make(chan anyThing)
	out2 := make(chan anyThing)
	go func() {
		for i := range inp {
			select { // send to whomever is ready to receive
			case out1 <- i:
			case out2 <- i:
			}
		}
	}()
	return out1, out2
}

// End of anyThingFork functions
// ===========================================================================
