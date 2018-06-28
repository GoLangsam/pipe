// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingPipe functions

// anyThingPipeFunc returns a channel to receive
// every result of action `act` applied to `inp`
// before close.
// Note: it 'could' be anyThingPipeMap for functional people,
// but 'map' has a very different meaning in go lang.
func anyThingPipeFunc(inp <-chan anyThing, act func(a anyThing) anyThing) (out <-chan anyThing) {
	cha := make(chan anyThing)
	if act == nil { // Make `nil` value useful
		act = func(a anyThing) anyThing { return a }
	}
	go func(out chan<- anyThing, inp <-chan anyThing, act func(a anyThing) anyThing) {
		defer close(out)
		for i := range inp {
			out <- act(i) // apply action
		}
	}(cha, inp, act)
	return cha
}

// End of anyThingPipe functions
// ===========================================================================
