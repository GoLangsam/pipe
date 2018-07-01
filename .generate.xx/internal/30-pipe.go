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
func anyThingPipe(inp anymode, ops ...func(a anyThing)) (out anymode) {
	cha := anymodeMakeChan()
	go pipeanyThing(cha, inp, ops...)
	return cha
}

func pipeanyThing(out anymode, inp anymode, ops ...func(a anyThing)) {
	defer out.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		for _, op := range ops {
			if op != nil {
				op(i) // chain action
			}
		}
		out.Provide(i) // send it
	}
}

// anyThingPipeFunc
// will chain every `act` to every `inp` and
// returns a channel to receive
// each result
// before close.
func anyThingPipeFunc(inp anymode, acts ...func(a anyThing) anyThing) (out anymode) {
	cha := anymodeMakeChan()
	go pipeanyThingFunc(cha, inp, acts...)
	return cha
}

func pipeanyThingFunc(out anymode, inp anymode, acts ...func(a anyThing) anyThing) {
	defer out.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		for _, act := range acts {
			if act != nil {
				i = act(i) // chain action
			}
		}
		out.Provide(i) // send result
	}
}

// End of anyThingPipe functions
// ===========================================================================
