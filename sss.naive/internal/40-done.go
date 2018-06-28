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
func anyThingDone(inp chan anyThing) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			_ = i // drain inp
		}
		done <- struct{}{}
	}()
	return done
}

// anyThingDoneSlice returns a channel to receive
// a slice with every anyThing received on `inp`
// upon close.
//
//  Note: Unlike anyThingDone, anyThingDoneSlice sends the fully accumulated slice, not just an event, once upon close of inp.
func anyThingDoneSlice(inp chan anyThing) chan []anyThing {
	done := make(chan []anyThing)
	go func() {
		defer close(done)
		slice := []anyThing{}
		for i := range inp {
			slice = append(slice, i)
		}
		done <- slice
	}()
	return done
}

// anyThingDoneFunc
// will apply `act` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func anyThingDoneFunc(inp chan anyThing, act func(a anyThing)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a anyThing) { return }
	}
	go func() {
		defer close(done)
		for i := range inp {
			act(i) // apply action
		}
		done <- struct{}{}
	}()
	return done
}

// End of anyThingDone terminators
// ===========================================================================
