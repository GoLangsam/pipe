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
func anyThingPipeFunc(inp Anymode, act func(a anyThing) anyThing) (out Anymode) {
	cha := MakeAnymodeChan()
	if act == nil {
		act = func(a anyThing) anyThing { return a }
	}
	go pipeanyThingFunc(cha, inp, act)
	return cha
}

func pipeanyThingFunc(out Anymode, inp Anymode, act func(a anyThing) anyThing) {
	defer out.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out.Provide(act(i))
	}
}

// anyThingPipeBuffer returns a buffered channel with capacity `cap` to receive
// all `inp`
// before close.
func anyThingPipeBuffer(inp Anymode, cap int) (out Anymode) {
	cha := MakeAnymodeBuff(cap)
	go pipeanyThingBuffer(cha, inp)
	return cha
}

func pipeanyThingBuffer(out Anymode, inp Anymode) {
	defer out.Close()
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		out.Provide(i)
	}
}

// End of anyThingPipe functions
// ===========================================================================
