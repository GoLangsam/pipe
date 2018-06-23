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
	go func() {
		for i := range inp {
			out <- act(i) // apply action
		}
	}()
	return out
}

// End of anyThingPipe functions
// ===========================================================================
