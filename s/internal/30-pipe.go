// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingPipe functions

// anyThingPipe
// will apply every `op` to every `inp` and
// returns a channel to receive
// each `inp`
// before close.
//
// Note: For functional people,
// this 'could' be named `anyThingMap`.
// Just: 'map' has a very different meaning in go lang.
func anyThingPipe(inp <-chan anyThing, ops ...func(a anyThing)) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go pipeanyThing(cha, inp, ops...)
	return cha
}

func pipeanyThing(out chan<- anyThing, inp <-chan anyThing, ops ...func(a anyThing)) {
	defer close(out)
	for i := range inp {
		for _, op := range ops {
			if op != nil {
				op(i) // chain action
			}
		}
		out <- i // send it
	}
}

// anyThingPipeFunc
// will chain every `act` to every `inp` and
// returns a channel to receive
// each result
// before close.
func anyThingPipeFunc(inp <-chan anyThing, acts ...func(a anyThing) anyThing) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go pipeanyThingFunc(cha, inp, acts...)
	return cha
}

func pipeanyThingFunc(out chan<- anyThing, inp <-chan anyThing, acts ...func(a anyThing) anyThing) {
	defer close(out)
	for i := range inp {
		for _, act := range acts {
			if act != nil {
				i = act(i) // chain action
			}
		}
		out <- i // send result
	}
}

// End of anyThingPipe functions
// ===========================================================================
