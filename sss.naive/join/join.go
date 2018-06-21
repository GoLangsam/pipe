// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// ===========================================================================
// Beg of anyThingJoin feedback back-feeders for circular networks

// anyThingJoin sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func anyThingJoin(out chan anyThing, inp ...anyThing) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}()
	return done
}

// anyThingJoinSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func anyThingJoinSlice(out chan anyThing, inp ...[]anyThing) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
		done <- struct{}{}
	}()
	return done
}

// anyThingJoinChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func anyThingJoinChan(out chan anyThing, inp chan anyThing) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// End of anyThingJoin feedback back-feeders for circular networks
// ===========================================================================
