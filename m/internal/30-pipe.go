// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of Pipe functions

// PipeFunc returns a channel to receive
// every result of action `act` applied to `inp`
// before close.
// Note: it 'could' be PipeMap for functional people,
// but 'map' has a very different meaning in go lang.
func (inp anyThingFrom) PipeFunc(act func(a anyThing) anyThing) (out anyThingFrom) {
	cha := make(chan anyThing)
	if act == nil { // Make `nil` value useful
		act = func(a anyThing) anyThing { return a }
	}
	go inp.pipeFunc(cha, act)
	return cha
}

func (inp anyThingFrom) pipeFunc(out anyThingInto, act func(a anyThing) anyThing) {
	defer close(out)
	for i := range inp {
		out <- act(i) // apply action
	}
}

// End of Pipe functions
// ===========================================================================
