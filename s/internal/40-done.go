// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingDone terminators

// anyThingDone returns a channel to receive
// one signal before close after `inp` has been drained.
func anyThingDone(inp <-chan anyThing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneanyThing(sig, inp)
	return sig
}

func doneanyThing(done chan<- struct{}, inp <-chan anyThing) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// anyThingDoneSlice returns a channel to receive
// a slice with every anyThing received on `inp`
// before close.
//
//  Note: Unlike anyThingDone, DoneanyThingSlice sends the fully accumulated slice, not just an event, once upon close of inp.
func anyThingDoneSlice(inp <-chan anyThing) (done <-chan []anyThing) {
	sig := make(chan []anyThing)
	go doneanyThingSlice(sig, inp)
	return sig
}

func doneanyThingSlice(done chan<- []anyThing, inp <-chan anyThing) {
	defer close(done)
	slice := []anyThing{}
	for i := range inp {
		slice = append(slice, i)
	}
	done <- slice
}

// anyThingDoneFunc returns a channel to receive
// one signal after `act` has been applied to every `inp`
// before close.
func anyThingDoneFunc(inp <-chan anyThing, act func(a anyThing)) (done <-chan struct{}) {
	sig := make(chan struct{})
	if act == nil {
		act = func(a anyThing) { return }
	}
	go doneanyThingFunc(sig, inp, act)
	return sig
}

func doneanyThingFunc(done chan<- struct{}, inp <-chan anyThing, act func(a anyThing)) {
	defer close(done)
	for i := range inp {
		act(i) // apply action
	}
	done <- struct{}{}
}

// End of anyThingDone terminators
// ===========================================================================
