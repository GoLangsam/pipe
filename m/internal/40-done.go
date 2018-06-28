// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of anyThingDone terminators

// anyThingDone returns a channel to receive
// one signal
// upon close
// and after `inp` has been drained.
func (inp anyThingFrom) anyThingDone() (done <-chan struct{}) {
	sig := make(chan struct{})
	go inp.doneanyThing(sig)
	return sig
}

func (inp anyThingFrom) doneanyThing(done chan<- struct{}) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// anyThingDoneSlice returns a channel to receive
// a slice with every anyThing received on `inp`
// upon close.
//
//  Note: Unlike anyThingDone, anyThingDoneSlice sends the fully accumulated slice, not just an event, once upon close of inp.
func (inp anyThingFrom) anyThingDoneSlice() (done <-chan []anyThing) {
	sig := make(chan []anyThing)
	go inp.doneanyThingSlice(sig)
	return sig
}

func (inp anyThingFrom) doneanyThingSlice(done chan<- []anyThing) {
	defer close(done)
	slice := []anyThing{}
	for i := range inp {
		slice = append(slice, i)
	}
	done <- slice
}

// anyThingDoneFunc
// will apply `act` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func (inp anyThingFrom) anyThingDoneFunc(act func(a anyThing)) (done <-chan struct{}) {
	sig := make(chan struct{})
	if act == nil {
		act = func(a anyThing) { return }
	}
	go inp.doneanyThingFunc(sig, act)
	return sig
}

func (inp anyThingFrom) doneanyThingFunc(done chan<- struct{}, act func(a anyThing)) {
	defer close(done)
	for i := range inp {
		act(i) // apply action
	}
	done <- struct{}{}
}

// End of anyThingDone terminators
// ===========================================================================
