// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingDone terminators

// anyThingDone returns a channel to receive
// one signal before close after `inp` has been drained.
func anyThingDone(inp anymode) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doitanyThing(sig, inp)
	return sig
}

func doitanyThing(done chan<- struct{}, inp anymode) {
	defer close(done)
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// anyThingDoneSlice returns a channel to receive
// a slice with every anyThing received on `inp`
// before close.
//
//  Note: Unlike anyThingDone, anyThingDoneSlice sends the fully accumulated slice, not just an event, once upon close of inp.
func anyThingDoneSlice(inp anymode) (done <-chan []anyThing) {
	sig := make(chan []anyThing)
	go doitanyThingSlice(sig, inp)
	return sig
}

func doitanyThingSlice(done chan<- []anyThing, inp anymode) {
	defer close(done)
	slice := []anyThing{}
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		slice = append(slice, i)
	}
	done <- slice
}

// anyThingDoneFunc returns a channel to receive
// one signal after `act` has been applied to every `inp`
// before close.
func anyThingDoneFunc(inp anymode, act func(a anyThing)) (done <-chan struct{}) {
	sig := make(chan struct{})
	if act == nil {
		act = func(a anyThing) { return }
	}
	go doitanyThingFunc(sig, inp, act)
	return sig
}

func doitanyThingFunc(done chan<- struct{}, inp anymode, act func(a anyThing)) {
	defer close(done)
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		act(i) // apply action
	}
	done <- struct{}{}
}

// End of anyThingDone terminators
// ===========================================================================
