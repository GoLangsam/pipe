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
func anyThingPipeFunc(inp anymode, act func(a anyThing) anyThing) (out anymode) {
	cha := anymodeMakeChan()
	if act == nil { // Make `nil` value useful
		act = func(a anyThing) anyThing { return a }
	}
	go pipeanyThingFunc(cha, inp, act)
	return cha
}

func pipeanyThingFunc(out anymode, inp anymode, act func(a anyThing) anyThing) {
	defer out.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out.Provide(act(i))
	}
}

// End of anyThingPipe functions
// ===========================================================================
