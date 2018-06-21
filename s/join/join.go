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
func anyThingJoin(out chan<- anyThing, inp ...anyThing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go joinanyThing(sig, out, inp...)
	return sig
}

func joinanyThing(done chan<- struct{}, out chan<- anyThing, inp ...anyThing) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// anyThingJoinSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func anyThingJoinSlice(out chan<- anyThing, inp ...[]anyThing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go joinanyThingSlice(sig, out, inp...)
	return sig
}

func joinanyThingSlice(done chan<- struct{}, out chan<- anyThing, inp ...[]anyThing) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// anyThingJoinChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func anyThingJoinChan(out chan<- anyThing, inp <-chan anyThing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go joinanyThingChan(sig, out, inp)
	return sig
}

func joinanyThingChan(done chan<- struct{}, out chan<- anyThing, inp <-chan anyThing) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// End of anyThingJoin feedback back-feeders for circular networks
// ===========================================================================
