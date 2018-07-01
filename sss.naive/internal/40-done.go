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
func anyThingDone(inp chan anyThing, ops ...func(a anyThing)) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			for _, op := range ops {
				if op != nil {
					op(i) // apply operation
				}
			}
		}
		done <- struct{}{}
	}()
	return done
}

// anyThingDoneFunc
// will chain every `act` to every `inp` and
// returns a channel to receive
// one signal
// upon close.
func anyThingDoneFunc(inp chan anyThing, acts ...func(a anyThing) anyThing) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			for _, act := range acts {
				if act != nil {
					i = act(i) // chain action
				}
			}
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

// End of anyThingDone terminators
// ===========================================================================
