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
// Beg of anyThingPipeEnter/Leave - Flapdoors observed by a Waiter

// anyThingWaiter - as implemented by `*sync.WaitGroup` -
// attends Flapdoors and keeps counting
// who enters and who leaves.
//
// Use DoneanyThingWait to learn about
// when the facilities are closed.
//
// Note: You may also use Your provided `*sync.WaitGroup.Wait()`
// to know when to close the facilities.
// Just: DoneanyThingWait is more convenient
// as it also closes the primary channel.
//
// Just make sure to have _all_ entrances and exits attended,
// and `Wait()` only *after* You've started flooding the facilities.
type anyThingWaiter interface {
	Add(delta int)
	Done()
	Wait()
}

// Note: Name is generic in order to avoid multiple-declaration clashes.

// anyThingPipeEnter returns a channel to receive
// all `inp`
// and registers throughput
// as arrival
// on the given `sync.WaitGroup`
// until close.
func anyThingPipeEnter(inp <-chan anyThing, wg anyThingWaiter) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go pipeanyThingEnter(cha, wg, inp)
	return cha
}

// anyThingPipeLeave returns a channel to receive
// all `inp`
// and registers throughput
// as departure
// on the given `sync.WaitGroup`
// until close.
func anyThingPipeLeave(inp <-chan anyThing, wg anyThingWaiter) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go pipeanyThingLeave(cha, wg, inp)
	return cha
}

func pipeanyThingEnter(out chan<- anyThing, wg anyThingWaiter, inp <-chan anyThing) {
	defer close(out)
	for i := range inp {
		wg.Add(1)
		out <- i
	}
}

func pipeanyThingLeave(out chan<- anyThing, wg anyThingWaiter, inp <-chan anyThing) {
	defer close(out)
	for i := range inp {
		out <- i
		wg.Done()
	}
}

// anyThingTubeEnter returns a closure around anyThingPipeEnter (_, wg)
// registering throughput
// on the given `sync.WaitGroup`
// as arrival.
func anyThingTubeEnter(wg anyThingWaiter) (tube func(inp <-chan anyThing) (out <-chan anyThing)) {

	return func(inp <-chan anyThing) (out <-chan anyThing) {
		return anyThingPipeEnter(inp, wg)
	}
}

// anyThingTubeLeave returns a closure around anyThingPipeLeave (_, wg)
// registering throughput
// on the given `sync.WaitGroup`
// as departure.
func anyThingTubeLeave(wg anyThingWaiter) (tube func(inp <-chan anyThing) (out <-chan anyThing)) {

	return func(inp <-chan anyThing) (out <-chan anyThing) {
		return anyThingPipeLeave(inp, wg)
	}
}

// anyThingDoneWait returns a channel to receive
// one signal
// after wg.Wait() has returned and inp has been closed
// before close.
//
// Note: Use only *after* You've started flooding the facilities.
func anyThingDoneWait(inp chan<- anyThing, wg anyThingWaiter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doneanyThingWait(cha, inp, wg)
	return cha
}

func doneanyThingWait(done chan<- struct{}, inp chan<- anyThing, wg anyThingWaiter) {
	defer close(done)
	wg.Wait()
	close(inp)
	done <- struct{}{} // not really needed - but looks better
}

// anyThingFiniWait returns a closure around `DoneanyThingWait(_, wg)`.
func anyThingFiniWait(wg anyThingWaiter) func(inp chan<- anyThing) (done <-chan struct{}) {

	return func(inp chan<- anyThing) (done <-chan struct{}) {
		return anyThingDoneWait(inp, wg)
	}
}

// End of anyThingPipeEnter/Leave - Flapdoors observed by a Waiter
// ===========================================================================
