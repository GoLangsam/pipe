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
// Use anyThingDoneWait to learn about
// when the facilities are closed.
//
// Note: You may also use Your provided `*sync.WaitGroup.Wait()`
// to know when to close the facilities.
// Just: anyThingDoneWait is more convenient
// as it also closes the primary channel for You.
//
// Just make sure to have _all_ entrances and exits attended,
// and `Wait()` only *after* You've started flooding the facilities.
type anyThingWaiter interface {
	Add(delta int)
	Done()
	Wait()
}

// Note: The name is intentionally generic in order to avoid eventual multiple-declaration clashes.

// anyThingPipeEnter returns a channel to receive
// all `inp`
// and registers throughput
// as arrival
// on the given `sync.WaitGroup`
// until close.
func (inp anyThingFrom) anyThingPipeEnter(wg anyThingWaiter) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go inp.pipeanyThingEnter(cha, wg)
	return cha
}

// anyThingPipeLeave returns a channel to receive
// all `inp`
// and registers throughput
// as departure
// on the given `sync.WaitGroup`
// until close.
func (inp anyThingFrom) anyThingPipeLeave(wg anyThingWaiter) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go inp.pipeanyThingLeave(cha, wg)
	return cha
}

// anyThingDoneLeave returns a channel to receive
// one signal after
// all throughput on `inp`
// has been registered
// as departure
// on the given `sync.WaitGroup`
// before close.
func (inp anyThingFrom) anyThingDoneLeave(wg anyThingWaiter) (done <-chan struct{}) {
	sig := make(chan struct{})
	go inp.doneanyThingLeave(sig, wg)
	return sig
}

func (inp anyThingFrom) pipeanyThingEnter(out anyThingInto, wg anyThingWaiter) {
	defer close(out)
	for i := range inp {
		wg.Add(1)
		out <- i
	}
}

func (inp anyThingFrom) pipeanyThingLeave(out anyThingInto, wg anyThingWaiter) {
	defer close(out)
	for i := range inp {
		out <- i
		wg.Done()
	}
}

func (inp anyThingFrom) doneanyThingLeave(done chan<- struct{}, wg anyThingWaiter) {
	defer close(done)
	for i := range inp {
		_ = i // discard
		wg.Done()
	}
	done <- struct{}{}
}

// anyThingTubeEnter returns a closure around anyThingPipeEnter (wg)
// registering throughput
// as arrival
// on the given `sync.WaitGroup`.
func (inp anyThingFrom) anyThingTubeEnter(wg anyThingWaiter) (tube func(inp anyThingFrom) (out anyThingFrom)) {

	return func(inp anyThingFrom) (out anyThingFrom) {
		return inp.anyThingPipeEnter(wg)
	}
}

// anyThingTubeLeave returns a closure around anyThingPipeLeave (wg)
// registering throughput
// as departure
// on the given `sync.WaitGroup`.
func (inp anyThingFrom) anyThingTubeLeave(wg anyThingWaiter) (tube func(inp anyThingFrom) (out anyThingFrom)) {

	return func(inp anyThingFrom) (out anyThingFrom) {
		return inp.anyThingPipeLeave(wg)
	}
}

// anyThingFiniLeave returns a closure around `anyThingDoneLeave(wg)`
// registering throughput
// as departure
// on the given `sync.WaitGroup`.
func (inp anyThingFrom) anyThingFiniLeave(wg anyThingWaiter) func(inp anyThingFrom) (done <-chan struct{}) {

	return func(inp anyThingFrom) (done <-chan struct{}) {
		return inp.anyThingDoneLeave(wg)
	}
}

// anyThingDoneWait returns a channel to receive
// one signal
// after wg.Wait() has returned and inp has been closed
// before close.
//
// Note: Use only *after* You've started flooding the facilities.
func (inp anyThingInto) anyThingDoneWait(wg anyThingWaiter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go inp.doneanyThingWait(cha, wg)
	return cha
}

func (inp anyThingInto) doneanyThingWait(done chan<- struct{}, wg anyThingWaiter) {
	defer close(done)
	wg.Wait()
	close(inp)
	done <- struct{}{} // not really needed - but looks better
}

// anyThingFiniWait returns a closure around `DoneanyThingWait(wg)`.
func (inp anyThingInto) anyThingFiniWait(wg anyThingWaiter) func(inp anyThingInto) (done <-chan struct{}) {

	return func(inp anyThingInto) (done <-chan struct{}) {
		return inp.anyThingDoneWait(wg)
	}
}

// End of anyThingPipeEnter/Leave - Flapdoors observed by a Waiter
// ===========================================================================
