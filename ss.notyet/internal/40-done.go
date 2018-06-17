// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of DoneAny terminators

// DoneAny returns a channel to receive
// one signal before close after `inp` has been drained.
func DoneAny(inp <-chan Any) (done <-chan struct{}) {
	sig := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan Any) {
		defer close(done)
		for i := range inp {
			_ = i // drain inp
		}
		done <- struct{}{}
	}(sig, inp)
	return sig
}

// DoneAnySlice returns a channel to receive
// a slice with every Any received on `inp`
// before close.
//
//  Note: Unlike DoneAny, DoneAnySlice sends the fully accumulated slice, not just an event, once upon close of inp.
func DoneAnySlice(inp <-chan Any) (done <-chan []Any) {
	sig := make(chan []Any)
	go func(done chan<- []Any, inp <-chan Any) {
		defer close(done)
		slice := []Any{}
		for i := range inp {
			slice = append(slice, i)
		}
		done <- slice
	}(sig, inp)
	return sig
}

// DoneAnyFunc returns a channel to receive
// one signal after `act` has been applied to every `inp`
// before close.
func DoneAnyFunc(inp <-chan Any, act func(a Any)) (done <-chan struct{}) {
	sig := make(chan struct{})
	if act == nil {
		act = func(a Any) { return }
	}
	go func(done chan<- struct{}, inp <-chan Any, act func(a Any)) {
		defer close(done)
		for i := range inp {
			act(i) // apply action
		}
		done <- struct{}{}
	}(sig, inp, act)
	return sig
}

// End of DoneAny terminators
// ===========================================================================
