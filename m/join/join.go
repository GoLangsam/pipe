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
// Beg of Join feedback back-feeders for circular networks

// Join sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func (out anyThingInto) Join(inp ...anyThing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go out.join(sig, inp...)
	return sig
}

func (out anyThingInto) join(done chan<- struct{}, inp ...anyThing) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func (out anyThingInto) JoinSlice(inp ...[]anyThing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go out.joinSlice(sig, inp...)
	return sig
}

func (out anyThingInto) joinSlice(done chan<- struct{}, inp ...[]anyThing) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func (out anyThingInto) JoinChan(inp anyThingFrom) (done <-chan struct{}) {
	sig := make(chan struct{})
	go out.joinChan(sig, inp)
	return sig
}

func (out anyThingInto) joinChan(done chan<- struct{}, inp anyThingFrom) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// End of Join feedback back-feeders for circular networks
// ===========================================================================
