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
func (out anyThingInto) anyThingJoin(inp ...anyThing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go out.joinanyThing(sig, inp...)
	return sig
}

func (out anyThingInto) joinanyThing(done chan<- struct{}, inp ...anyThing) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// anyThingJoinSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func (out anyThingInto) anyThingJoinSlice(inp ...[]anyThing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go out.joinanyThingSlice(sig, inp...)
	return sig
}

func (out anyThingInto) joinanyThingSlice(done chan<- struct{}, inp ...[]anyThing) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// anyThingJoinChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func (out anyThingInto) anyThingJoinChan(inp anyThingFrom) (done <-chan struct{}) {
	sig := make(chan struct{})
	go out.joinanyThingChan(sig, inp)
	return sig
}

func (out anyThingInto) joinanyThingChan(done chan<- struct{}, inp anyThingFrom) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// End of anyThingJoin feedback back-feeders for circular networks
// ===========================================================================
