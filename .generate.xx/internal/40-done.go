// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of DoneAny terminators

// DoneAny returns a channel to receive
// one signal before close after `inp` has been drained.
func DoneAny(inp Anymode) (done <-chan struct{}) {
	sig := make(chan struct{})
	go doitAny(sig, inp)
	return sig
}

func doitAny(done chan<- struct{}, inp Anymode) {
	defer close(done)
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneAnySlice returns a channel to receive
// a slice with every Any received on `inp`
// before close.
//
//  Note: Unlike DoneAny, DoneAnySlice sends the fully accumulated slice, not just an event, once upon close of inp.
func DoneAnySlice(inp Anymode) (done <-chan []Any) {
	sig := make(chan []Any)
	go doitAnySlice(sig, inp)
	return sig
}

func doitAnySlice(done chan<- []Any, inp Anymode) {
	defer close(done)
	slice := []Any{}
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		slice = append(slice, i)
	}
	done <- slice
}

// DoneAnyFunc returns a channel to receive
// one signal after `act` has been applied to every `inp`
// before close.
func DoneAnyFunc(inp Anymode, act func(a Any)) (done <-chan struct{}) {
	sig := make(chan struct{})
	if act == nil {
		act = func(a Any) { return }
	}
	go doitAnyFunc(sig, inp, act)
	return sig
}

func doitAnyFunc(done chan<- struct{}, inp Anymode, act func(a Any)) {
	defer close(done)
	for i, ok := inp.Request(); ok; i, ok = inp.Request() {
		act(i) // apply action
	}
	done <- struct{}{}
}

// End of DoneAny terminators
// ===========================================================================
