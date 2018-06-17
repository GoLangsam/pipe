// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

type Any generic.Type

// ===========================================================================
// Beg of JoinAny feedback back-feeders for circular networks

// JoinAny sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinAny(out chan<- Any, inp ...Any) (done <-chan struct{}) {
	sig := make(chan struct{})
	go joinAny(sig, out, inp...)
	return sig
}

func joinAny(done chan<- struct{}, out chan<- Any, inp ...Any) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinAnySlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinAnySlice(out chan<- Any, inp ...[]Any) (done <-chan struct{}) {
	sig := make(chan struct{})
	go joinAnySlice(sig, out, inp...)
	return sig
}

func joinAnySlice(done chan<- struct{}, out chan<- Any, inp ...[]Any) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinAnyChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinAnyChan(out chan<- Any, inp <-chan Any) (done <-chan struct{}) {
	sig := make(chan struct{})
	go joinAnyChan(sig, out, inp)
	return sig
}

func joinAnyChan(done chan<- struct{}, out chan<- Any, inp <-chan Any) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// End of JoinAny feedback back-feeders for circular networks
// ===========================================================================
