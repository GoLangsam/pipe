// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingPipe functions

// anyThingPipeFunc returns a channel to receive
// every result of action `act` applied to `inp`
// before close.
//  Note: it 'could' be anyThingPipeMap for functional people,
//  but 'map' has a very different meaning in go lang.
func anyThingPipeFunc(inp chan anyThing, act func(a anyThing) anyThing) chan anyThing {
	out := make(chan anyThing)
	if act == nil { // Make `nil` value useful
		act = func(a anyThing) anyThing { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i) // apply action
		}
	}()
	return out
}

// anyThingPipeBuffer returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func anyThingPipeBuffer(inp chan anyThing, cap int) chan anyThing {
	out := make(chan anyThing, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// End of anyThingPipe functions
// ===========================================================================
