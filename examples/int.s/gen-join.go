// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

// ===========================================================================
// Beg of intJoin feedback back-feeders for circular networks

// intJoin sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func intJoin(out chan<- int, inp ...int) (done <-chan struct{}) {
	sig := make(chan struct{})
	go joinint(sig, out, inp...)
	return sig
}

func joinint(done chan<- struct{}, out chan<- int, inp ...int) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// intJoinSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func intJoinSlice(out chan<- int, inp ...[]int) (done <-chan struct{}) {
	sig := make(chan struct{})
	go joinintSlice(sig, out, inp...)
	return sig
}

func joinintSlice(done chan<- struct{}, out chan<- int, inp ...[]int) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// intJoinChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func intJoinChan(out chan<- int, inp <-chan int) (done <-chan struct{}) {
	sig := make(chan struct{})
	go joinintChan(sig, out, inp)
	return sig
}

func joinintChan(done chan<- struct{}, out chan<- int, inp <-chan int) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// End of intJoin feedback back-feeders for circular networks
// ===========================================================================
