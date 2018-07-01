// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingDone terminators

// anyThingDone
// will apply every `op` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func anyThingDone(inp anymode, ops ...func(a anyThing)) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneanyThing(sig, inp, ops...)
	return sig
}

func doneanyThing(done chan<- struct{}, inp anymode, ops ...func(a anyThing)) {
	defer close(done)
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		for _, op := range ops {
			if op != nil {
				op(i) // apply operation
			}
		}
	}
	done <- struct{}{}
}

// anyThingDoneFunc
// will chain every `act` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func anyThingDoneFunc(inp anymode, acts ... func(a anyThing) anyThing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doneanyThingFunc(sig, inp, acts...)
	return sig
}

func doneanyThingFunc(done chan<- struct{}, inp anymode, acts ... func(a anyThing) anyThing) {
	defer close(done)
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		for _, act := range acts {
			if act != nil {
				i = act(i) // chain action
			}
		}
	}
	done <- struct{}{}
}

// anyThingDoneSlice returns a channel to receive
// a slice with every anyThing received on `inp`
// upon close.
//
//  Note: Unlike anyThingDone, anyThingDoneSlice sends the fully accumulated slice, not just an event, once upon close of inp.
func anyThingDoneSlice(inp anymode) (done <-chan []anyThing) {
	sig := make(chan []anyThing)
	go doneanyThingSlice(sig, inp)
	return sig
}

func doneanyThingSlice(done chan<- []anyThing, inp anymode) {
	defer close(done)
	slice := []anyThing{}
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		slice = append(slice, i)
	}
	done <- slice
}

// End of anyThingDone terminators
// ===========================================================================
