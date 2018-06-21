// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of ThingJoin feedback back-feeders for circular networks

// ThingJoin sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func ThingJoin(out chan<- Thing, inp ...Thing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go joinThing(sig, out, inp...)
	return sig
}

func joinThing(done chan<- struct{}, out chan<- Thing, inp ...Thing) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// ThingJoinSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func ThingJoinSlice(out chan<- Thing, inp ...[]Thing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go joinThingSlice(sig, out, inp...)
	return sig
}

func joinThingSlice(done chan<- struct{}, out chan<- Thing, inp ...[]Thing) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// ThingJoinChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func ThingJoinChan(out chan<- Thing, inp <-chan Thing) (done <-chan struct{}) {
	sig := make(chan struct{})
	go joinThingChan(sig, out, inp)
	return sig
}

func joinThingChan(done chan<- struct{}, out chan<- Thing, inp <-chan Thing) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// End of ThingJoin feedback back-feeders for circular networks
