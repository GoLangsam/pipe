// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pipe

import (
	"github.com/cheekybits/genny/generic"
)

// anyThing is the generic type flowing thru the pipe network.
type anyThing generic.Type

// anyOwner is the generic who shall own the methods.
//  Note: Need to use `generic.Number` here as `generic.Type` is an interface and cannot have any method.
type anyOwner generic.Number

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
func (my anyOwner) anyThingPipeEnter(inp <-chan anyThing, wg anyThingWaiter) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go my.pipeanyThingEnter(cha, wg, inp)
	return cha
}

// anyThingPipeLeave returns a channel to receive
// all `inp`
// and registers throughput
// as departure
// on the given `sync.WaitGroup`
// until close.
func (my anyOwner) anyThingPipeLeave(inp <-chan anyThing, wg anyThingWaiter) (out <-chan anyThing) {
	cha := make(chan anyThing)
	go my.pipeanyThingLeave(cha, wg, inp)
	return cha
}

// anyThingDoneLeave returns a channel to receive
// one signal after
// all throughput on `inp`
// has been registered
// as departure
// on the given `sync.WaitGroup`
// before close.
func (my anyOwner) anyThingDoneLeave(inp <-chan anyThing, wg anyThingWaiter) (done <-chan struct{}) {
	sig := make(chan struct{})
	go my.doneanyThingLeave(sig, wg, inp)
	return sig
}

func (my anyOwner) pipeanyThingEnter(out chan<- anyThing, wg anyThingWaiter, inp <-chan anyThing) {
	defer close(out)
	for i := range inp {
		wg.Add(1)
		out <- i
	}
}

func (my anyOwner) pipeanyThingLeave(out chan<- anyThing, wg anyThingWaiter, inp <-chan anyThing) {
	defer close(out)
	for i := range inp {
		out <- i
		wg.Done()
	}
}

func (my anyOwner) doneanyThingLeave(done chan<- struct{}, wg anyThingWaiter, inp <-chan anyThing) {
	defer close(done)
	for i := range inp {
		_ = i // discard
		wg.Done()
	}
	done <- struct{}{}
}

// anyThingTubeEnter returns a closure around anyThingPipeEnter (_, wg)
// registering throughput
// as arrival
// on the given `sync.WaitGroup`.
func (my anyOwner) anyThingTubeEnter(wg anyThingWaiter) (tube func(inp <-chan anyThing) (out <-chan anyThing)) {

	return func(inp <-chan anyThing) (out <-chan anyThing) {
		return my.anyThingPipeEnter(inp, wg)
	}
}

// anyThingTubeLeave returns a closure around anyThingPipeLeave (_, wg)
// registering throughput
// as departure
// on the given `sync.WaitGroup`.
func (my anyOwner) anyThingTubeLeave(wg anyThingWaiter) (tube func(inp <-chan anyThing) (out <-chan anyThing)) {

	return func(inp <-chan anyThing) (out <-chan anyThing) {
		return my.anyThingPipeLeave(inp, wg)
	}
}

// anyThingFiniLeave returns a closure around `anyThingDoneLeave(_, wg)`
// registering throughput
// as departure
// on the given `sync.WaitGroup`.
func (my anyOwner) anyThingFiniLeave(wg anyThingWaiter) func(inp <-chan anyThing) (done <-chan struct{}) {

	return func(inp <-chan anyThing) (done <-chan struct{}) {
		return my.anyThingDoneLeave(inp, wg)
	}
}

// anyThingDoneWait returns a channel to receive
// one signal
// after wg.Wait() has returned and inp has been closed
// before close.
//
// Note: Use only *after* You've started flooding the facilities.
func (my anyOwner) anyThingDoneWait(inp chan<- anyThing, wg anyThingWaiter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go my.doneanyThingWait(cha, inp, wg)
	return cha
}

func (my anyOwner) doneanyThingWait(done chan<- struct{}, inp chan<- anyThing, wg anyThingWaiter) {
	defer close(done)
	wg.Wait()
	close(inp)
	done <- struct{}{} // not really needed - but looks better
}

// anyThingFiniWait returns a closure around `DoneanyThingWait(_, wg)`.
func (my anyOwner) anyThingFiniWait(wg anyThingWaiter) func(inp chan<- anyThing) (done <-chan struct{}) {

	return func(inp chan<- anyThing) (done <-chan struct{}) {
		return my.anyThingDoneWait(inp, wg)
	}
}

// End of anyThingPipeEnter/Leave - Flapdoors observed by a Waiter
// ===========================================================================