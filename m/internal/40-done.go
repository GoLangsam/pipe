// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of Done terminators

// Done
// will apply every `op` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func (inp anyThingFrom) Done(ops ...func(a anyThing)) (done <-chan struct{}) {
	sig := make(chan struct{})
	go inp.done(sig, ops...)
	return sig
}

func (inp anyThingFrom) done(done chan<- struct{}, ops ...func(a anyThing)) {
	defer close(done)
	for i := range inp {
		for _, op := range ops {
			if op != nil {
				op(i) // apply operation
			}
		}
	}
	done <- struct{}{}
}

// DoneFunc
// will chain every `act` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func (inp anyThingFrom) DoneFunc(acts ...func(a anyThing) anyThing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go inp.doneFunc(sig, acts...)
	return sig
}

func (inp anyThingFrom) doneFunc(done chan<- struct{}, acts ...func(a anyThing) anyThing) {
	defer close(done)
	for i := range inp {
		for _, act := range acts {
			if act != nil {
				i = act(i) // chain action
			}
		}
	}
	done <- struct{}{}
}

// DoneSlice returns a channel to receive
// a slice with every anyThing received on `inp`
// upon close.
//
//  Note: Unlike Done, DoneSlice sends the fully accumulated slice, not just an event, once upon close of inp.
func (inp anyThingFrom) DoneSlice() (done <-chan []anyThing) {
	sig := make(chan []anyThing)
	go inp.doneSlice(sig)
	return sig
}

func (inp anyThingFrom) doneSlice(done chan<- []anyThing) {
	defer close(done)
	slice := []anyThing{}
	for i := range inp {
		slice = append(slice, i)
	}
	done <- slice
}

// End of Done terminators
// ===========================================================================
